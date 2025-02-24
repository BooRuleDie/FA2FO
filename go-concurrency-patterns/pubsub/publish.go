package pubsub

import (
	"context"
	"errors"
	"fmt"
	"log"
)

// newPublisher creates and returns a new Publisher instance with an initialized subscribers map.
// Other fields will be initialized with their zero values.
func newPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[string]*Subscriber),
		// rest of the fields will be initialized with
		// their zero values
	}
}

// subscribe adds a new subscriber with the given id and returns a receive-only channel for messages.
// Returns error if publisher is closed or if subscriber id already exists.
func (p *Publisher) subscribe(ctx context.Context, id string) (<-chan Message, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil, errors.New("can't subscribe to a closed publisher")
	}

	if _, ok := p.subscribers[id]; ok {
		return nil, errors.New("id already found in the subscribers list")
	}

	msgs := make(chan Message, 10)
	done := make(chan struct{})
	p.subscribers[id] = &Subscriber{
		id:   id,
		msgs: msgs,
		done: done,
	}

	// Start cleanup goroutine that unsubscribes when context is done or subscriber is removed
	go func() {
		select {
		case <-ctx.Done():
			p.unsubscribe(id)
			log.Println("ctx.Done() is called!")
		case <-done:
			// Cleanup is already handled by unsubscribe
		}
	}()

	return msgs, nil
}

// unsubscribe removes a subscriber with the given id.
// Returns error if publisher is closed or if subscriber id is not found.
func (p *Publisher) unsubscribe(id string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return errors.New("can't unsubscribe to a closed publisher")
	}

	sub, ok := p.subscribers[id]
	if !ok {
		return errors.New("id not found in the subscribers list")
	}

	// Close subscriber channels and remove from publisher's subscribers list
	close(sub.done) // Signal cleanup goroutine to exit
	close(sub.msgs)
	delete(p.subscribers, id)

	return nil
}

// publish sends a message to all current subscribers.
// Returns error if publisher is closed or if any subscriber's message channel is full.
func (p *Publisher) publish(msg Message) error {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.closed {
		return errors.New("can't publish to a closed publisher")
	}

	for _, sub := range p.subscribers {
		select {
		case sub.msgs <- msg:
		default:
			return fmt.Errorf("Warning: Dropping message for subscriber %s: channel full\n", sub.id)
		}
	}

	return nil
}

// close stops the publisher and removes all subscribers.
// Returns error if publisher is already closed.
func (p *Publisher) close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return errors.New("can't close a closed publisher")
	}

	p.closed = true
	for id, sub := range p.subscribers {
		close(sub.done)
		close(sub.msgs)
		delete(p.subscribers, id)
	}

	return nil
}
