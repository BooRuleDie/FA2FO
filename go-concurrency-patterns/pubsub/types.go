package pubsub

import (
	"sync"
)

type Message struct {
	topic   string
	content string
}

type Subscriber struct {
	id   string
	msgs chan Message
	done chan struct{}
}

type Publisher struct {
	mu          sync.RWMutex
	subscribers map[string]*Subscriber
	closed      bool
}

