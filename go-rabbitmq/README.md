# Usage
To interact with the system, use the following commands for publishing and consuming messages:
```bash
make publish
make consume
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

