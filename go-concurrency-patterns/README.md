### Covered Concurrency Patterns

- [x] **Worker Pools**: Distributing tasks across a fixed number of goroutines that process jobs from a shared queue
- [x] **Pipeline**: Connecting stages where each stage performs part of the overall task and passes results to the next  
- [ ] **Fan-out/Fan-in**: Distributing work across multiple goroutines and collecting results back into a single channel
- [ ] **Generator Pattern**: Using channels to generate a sequence of values
- [ ] **Pub/Sub**: Broadcasting messages to multiple subscribers through channels
- [ ] **Mutex and Read/Write Mutex**: Protecting shared resources from concurrent access
- [ ] **Context Pattern**: Managing cancellation, deadlines, and request-scoped values across API boundaries
- [ ] **Semaphore Pattern**: Limiting concurrent access to resources using buffered channels
- [ ] **Error Group**: Synchronizing multiple goroutines and collecting their errors
- [ ] **Select Pattern**: Coordinating multiple channels and handling timeouts

# Pipeline

In the pipeline pattern, a process is divided into stages, where each stage is a function with an upstream and a downstream channel. A stage takes data from its upstream, processes it, and sends the result to its downstream. The next stage then consumes this data as its upstream input, and this process continues until the last stage completes the pipeline.

The only exception to this flow is the first stage, also known as the producer or generator stage. Unlike other stages, it does not have an upstream channel—only a downstream—since it is responsible for initiating the pipeline. For further details, you can inspect the implementation in the following file: `pipeline/pipeline.go`.

# Worker Pool

They're very similar to thread pools. Essentially, you have a couple of goroutines that are always listening to specific channels, such as the jobs and results channels. They continuously consume from the jobs channel, process the data, and then send it to the results channel. The name of the pattern comes from the number of goroutines involved. 

In this pattern, there are generally multiple goroutines listening to the same channel, which makes them resemble workers. For further details, you can inspect the implementation in this file: `workerpools/workerpools.go`.