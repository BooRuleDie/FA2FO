# Usage

To interact with the system, use the following commands for publishing and consuming messages:

```bash
make publish
make consume

make publish-direct info|warning|error "some logs here"
make consume-direct info|warning|error|foo|bar...

make publish-topic kern "some kernel error" # <facility> <log>
make consume-topic kern.#|*.critical|*.error...

make rpc-server
make rpc-client 1 2 3 4 5 6...
```

# Hello World

![Hello World](./img/hello-world.png)

As shown in the illustration, this represents a basic implementation with a single publisher, consumer, and queue. As the name suggests, this is a "Hello World" application for the RabbitMQ message broker that demonstrates how to set up a queue, publish messages, and consume messages from the queue. To examine the implementation details, please review the contents of the `workerqueues` directory.

# Worker Queues

![Work Queues](./img/work-queues.png)

The same template is used as the "Hello World" example, but some configurations are changed. For instance, messages are durable in this example. In order to make a message durable in RabbitMQ, you need to set the durable flag to true in queue definition and also set the delivery mode as Persistent when publishing messages.

```go
// consumer.go
q, err := ch.QueueDeclare(
	"tasks", // name
	true,    // DURABLE <--
	false,   // delete when unused
	false,   // exclusive
	false,   // no-wait
	nil,     // arguments
)
```

```go
// publisher.go
err = ch.PublishWithContext(
	ctx,
	"",     // exchange
	q.Name, // routing key
	false,  // mandatory
	false,  // immediate
	amqp.Publishing{
		DeliveryMode: amqp.Persistent, // <--
		ContentType: "text/plain",
		Body:        []byte(body),
	},
)
```

It's important to mention that RabbitMQ doesn't use `fsync` like relational databases when you commit the transaction, which means there is a chance that the message is written to the cache instead of the disk. If the system crashes, you might still lose data, but adjusting this configuration significantly increases your system's reliability.

Also, auto-acknowledgment is disabled so that if a consumer fails to consume a task, it can be re-queued for another consumer.

```go
// consumer.go
msgs, err := ch.Consume(
	q.Name, // queue
	"",     // consumer
	false,  // AUTO-ACK <--
	false,  // exclusive
	false,  // no-local
	false,  // no-wait
	nil,    // args
)
```

And lastly, a QoS (Quality of Service) configuration is added for fair dispatch.

```go
ch.Qos(
	1,     // prefetchCount
	0,     // prefetchSize
	false, // global
)
```

The configuration basically states that the maximum amount of unacknowledged messages you can send to a consumer is 1, which means RabbitMQ can't send more than one message unless the previous message it sent to the consumer is acknowledged. `prefetchSize` does the same thing but instead of counting the number of messages sent, it looks at the size of the messages. When you give 0, RabbitMQ ignores this value. And if there is more than one consumer per channel, the global parameter sets the configuration for all of them. In our example, there is just one consumer, so it's set to false.

# Publish / Subscribe

![Publish / Subscribe](./img/pub-sub.png)

So far we've bound queues with consumers and publishers directly; however, this is not the case in most applications. Most of the time, the producer doesn't know which consumers are going to consume the message that's sent to the broker. It just sends messages to an Exchange and it's the Exchange's job to know how to route the messages to the right queues. If you inspect the `publisher.go` file, you can see that there is no queue declaration but only an Exchange declaration, and when publishing the messages, we've removed the old routing key which was a queue name and instead put an Exchange name there.

There are multiple Exchange types like direct, topic, fanout... The fanout type is used here, which sends the same message to all bound queues.

There are also some changes in the `consumer.go` file. In the queue definition, there is no name for the queue; this tells RabbitMQ to generate one for us. Also, the exclusive flag is set to true. Exclusive makes a queue private, which means only the current connection can access the queue, only one consumer can access the queue, and when the connection is closed, the queue will be removed _(even though there are messages inside)_. So it can be considered as a temporary private queue for the consumer.

We've also declared the Exchange here, and after both Exchange and queue are declared, we bind them together so the Exchange knows which queues to send messages to. Auto-acknowledgment is enabled in this example.

# Routing

![Routing](./img/routing.png)

In this example, we explore the direct exchange type and sophisticated routing concepts. While the structure closely resembles the previous example, we've introduced a critical difference: when publishing messages, we explicitly specify a `routingKey` that corresponds to message severity (error, warning, info). On the consumer side, before listening on queues, we bind each consumer with one or more particular `routingKeys`. This binding instructs the exchange precisely which queue should receive each message based on its routing key. This selective message distribution allows consumers to process only the message types they're configured to handle, creating a more targeted and efficient messaging system.

```go
// publisher.go
err = ch.PublishWithContext(
		ctx,
		exchangeName, // exchange
		severity,     // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(logMessage),
		},
	)
```

```go
// consumer.go
sevs := getSeverities()
	for _, sev := range sevs {
		err = ch.QueueBind(
			q.Name,       // queue name
			sev,          // routing key
			exchangeName, // exchange
			false,
			nil,
		)
		failOnError(err, "Failed to bind a queue")
	}
```

```bash
# start consumers
make consume-direct error                   # consumer 1 only listens for errors
make consume-direct info warning # consumer 2 listens for both warnings and info logs

# publish
make publish-direct info "some info log"
make publish-direct error "some error log"
make publish-direct warning "some warning log"
```

# Topic

![Topic](./img/topic.png)

This example is very similar to the previous one; however, there's a slight difference which is the type of the exchange. The `topic` exchange is used in this example and it can be considered as an advanced direct exchange. In a direct exchange, you can just specify the routing key for publisher and binding key as string values; however, in a topic exchange, you can also specify wildcards. There are two types of wildcards that you can use:

1. `#`: Which means zero or more words
2. `*`: Which means any one word

So if you set up the consumer and publisher with the following commands:

```bash
# consumer 1
make consume-topic "\#.critical"

# consumer 2
make consume-topic "kern.\*" "*.error"

# publisher
make publish-topic kern.error "some error log"
make publish-topic user.error "some error log"
make publish-topic kern.critical "some error log"
```

The first two messages are consumed by consumer 2, but the last one is consumed by both consumers, thanks to the wildcards specified on the consumers.

# RPC
![Hello World](./img/rpc.png)

Our final example demonstrates the implementation of Remote Procedure Call (RPC) using RabbitMQ. In this pattern, both client and server function simultaneously as publishers and consumers, creating a bidirectional communication channel. 

The client follows these steps in sequence:
1. Collects arguments that require processing
2. Creates an exclusive queue dedicated to receiving responses
3. Generates a unique correlation ID to track the request-response lifecycle
4. Assembles a complete message containing the correlation ID, reply-to queue information, and processing arguments
5. Publishes this message to the designated RPC server queue

On the server side, the process unfolds as follows:
1. Establishes a dedicated RPC queue for incoming requests
2. Maintains an active listener on this queue to capture client requests
3. Processes each incoming message according to the requested operation
4. Returns the processed result to the client's reply-to queue, preserving the original correlation ID to ensure proper request-response matching

This elegant pattern demonstrates how RabbitMQ can be leveraged to implement a robust, asynchronous RPC mechanism in Go, enabling distributed processing while maintaining clear request-response relationships.

