# Pipeline

In the pipeline pattern, a process is divided into stages, where each stage is a function with an upstream and a downstream channel. A stage takes data from its upstream, processes it, and sends the result to its downstream. The next stage then consumes this data as its upstream input, and this process continues until the last stage completes the pipeline.

The only exception to this flow is the first stage, also known as the producer or generator stage. Unlike other stages, it does not have an upstream channel—only a downstream—since it is responsible for initiating the pipeline. For further details, you can inspect the implementation in the following file: `pipeline/pipeline.go`.

# Worker Pool

They're very similar to thread pools. Essentially, you have a couple of goroutines that are always listening to specific channels, such as the jobs and results channels. They continuously consume from the jobs channel, process the data, and then send it to the results channel. The name of the pattern comes from the number of goroutines involved. 

In this pattern, there are generally multiple goroutines listening to the same channel, which makes them resemble workers. For further details, you can inspect the implementation in this file: `workerpools/workerpools.go`.