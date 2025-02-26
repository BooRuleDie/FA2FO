# Usage
To interact with the system, use the following commands for publishing and consuming messages:
```
make publish
make consume
```

# Hello World
![Hello World](./img/hello-world.png)

As shown in the illustration, this represents a basic implementation with a single publisher, consumer, and queue. As the name suggests, this is a "Hello World" application for the RabbitMQ message broker that demonstrates how to set up a queue, publish messages, and consume messages from the queue. To examine the implementation details, please review the contents of the `workerqueues` directory.

# Worker Queues
![Work Queues](./img/work-queues.png)

explain the changes:
* what to do for durable messages
    * queue def -> durable
    * deliveryMode of the messages -> persistent
* disable auto-ack
* Qos for fair dispatch
    * explain parameters