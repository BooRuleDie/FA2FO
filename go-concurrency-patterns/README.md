### Covered Concurrency Patterns

- [x] **Worker Pools**: Distributing tasks across a fixed number of goroutines that process jobs from a shared queue
- [x] **Pipeline**: Connecting stages where each stage performs part of the overall task and passes results to the next
- [x] **Fan-out/Fan-in**: Distributing work across multiple goroutines and collecting results back into a single channel
- [x] **Generator Pattern**: Using channels to generate a sequence of values
- [x] **Error Group**: Synchronizing multiple goroutines and collecting their errors
- [x] **Semaphore Pattern**: Limiting concurrent access to resources using buffered channels
- [x] **Mutex and Read/Write Mutex**: Protecting shared resources from concurrent access
- [x] **Pub/Sub**: Broadcasting messages to multiple subscribers through channels

# Pipeline

In the pipeline pattern, a process is divided into stages, where each stage is a function with an upstream and a downstream channel. A stage takes data from its upstream, processes it, and sends the result to its downstream. The next stage then consumes this data as its upstream input, and this process continues until the last stage completes the pipeline.

The only exception to this flow is the first stage, also known as the producer or generator stage. Unlike other stages, it does not have an upstream channel—only a downstream—since it is responsible for initiating the pipeline. For further details, you can inspect the implementation in the following file: `pipeline/pipeline.go`.

# Worker Pool

They're very similar to thread pools. Essentially, you have a couple of goroutines that are always listening to specific channels, such as the jobs and results channels. They continuously consume from the jobs channel, process the data, and then send it to the results channel. The name of the pattern comes from the number of goroutines involved.

In this pattern, there are generally multiple goroutines listening to the same channel, which makes them resemble workers. For further details, you can inspect the implementation in this file: `workerpools/workerpools.go`.

# Fan-in & Fan-out

The Fan-in/Fan-out pattern is a powerful concurrency design that combines two complementary operations: distributing work across multiple goroutines (fan-out) and consolidating results back into a single channel (fan-in). In the fan-out phase, a single channel's data is distributed to multiple goroutines for parallel processing. The fan-in phase then merges the output from these multiple goroutines back into a single channel, effectively combining their results.

This pattern is particularly useful for CPU-intensive tasks that can benefit from parallel processing, while maintaining ordered data flow through the system. Despite its simple concept, it's one of the most frequently used and effective concurrency patterns in Go. For a detailed implementation example, you can refer to the code in: `faninfanout/faninfanout.go`

# Generator

The Generator pattern has been utilized as a foundational component in several of the concurrency patterns we've implemented. It's a simple but powerful pattern that takes variadic input and returns a receive-only channel. A goroutine launched by the generator continuously sends values from the input to the channel, allowing other parts of the program to receive and process these values asynchronously. For implementation details, you can refer to the file: `generator/generator.go`

# Error Group

The Error Group pattern is a simple concurrency pattern that allows you to manage multiple goroutines and gather their errors. It uses the errgroup package, where a single error from any goroutine will cause `eg.Wait()` to return that error, signaling a failure. Wait blocks until all function calls from the Go method have returned, then returns the first non-nil error (if any) from them.

For a full implementation, check the `errgroup/errgroup.go` file.

# Semaphore

This pattern is all about rate limiting with the help of a buffered channel. The logic is simple - all worker goroutines send a signal to a buffered channel before starting work and also receive from the same channel after they're completed. If the buffered channel is full, the send operation blocks so the goroutine stops until at least one other worker goroutine receives from the buffered channel. For the implementation, you can inspect the `semaphore/semaphore.go` file.

# Mutex

Mutexes are used to lock shared data among goroutines so it can be updated or read without triggering a race condition. In Go, it's mostly recommended to use channels for communication across goroutines; however, that doesn't mean you should always be using channels for communication. When it comes to simpler tasks like increasing a counter, mutexes can be more appropriate. You can inspect the `mutex/mutex.go` file for further details.

# Pub / Sub

This is the most complicated pattern among all those covered so far. You can think of it as a message broker implementation in Go. There is a publisher and a consumer, and the rest of the operations are the same as message brokers. The tricky part is the usage of `RWMutex` in the publisher. If you inspect the code, you can see that when there is a write operation *(such as adding/removing subscribers, modifying subscription lists, updating shared state, closing channels, deleting subscribers, or changing publisher status)* to any shared data, `mu.Lock()` & `mu.Unlock()` are used, and when there's only a read operation, `mu.RLock()` & `mu.RUnlock()` are used. We can explain it with the following example:

* `mu.RLock()`: I'm going to do a read operation, so any other goroutine that accesses the same data as me shouldn't be able to write it until I'm finished; however, they can read the data.
* `mu.Lock()`: I'm going to do a write operation, so in order to prevent race conditions and data inconsistencies, no other goroutine should be able to write or read the same data I'm currently accessing.