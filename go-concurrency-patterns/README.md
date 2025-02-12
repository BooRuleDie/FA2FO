# Pipeline

In the pipeline pattern, a process is divided into stages, where each stage is a function with an upstream and a downstream channel. A stage takes data from its upstream, processes it, and sends the result to its downstream. The next stage then consumes this data as its upstream input, and this process continues until the last stage completes the pipeline.

The only exception to this flow is the first stage, also known as the producer or generator stage. Unlike other stages, it does not have an upstream channel—only a downstream—since it is responsible for initiating the pipeline.

For further details, you can inspect the implementation in the following file: `pipeline/pipeline.go`.