### Covered Concurrency Patterns

- [x] **Worker Pools**: Distributing tasks across a fixed number of goroutines that process jobs from a shared queue
- [x] **Pipeline**: Connecting stages where each stage performs part of the overall task and passes results to the next  
- [x] **Fan-out/Fan-in**: Distributing work across multiple goroutines and collecting results back into a single channel
- [x] **Generator Pattern**: Using channels to generate a sequence of values
- [x] **Error Group**: Synchronizing multiple goroutines and collecting their errors
- [ ] **Pub/Sub**: Broadcasting messages to multiple subscribers through channels
- [ ] **Mutex and Read/Write Mutex**: Protecting shared resources from concurrent access
- [ ] **Semaphore Pattern**: Limiting concurrent access to resources using buffered channels
- [ ] **Select Pattern**: Coordinating multiple channels and handling timeouts

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

