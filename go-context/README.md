# What are Contexts?

Contexts are data structures that help us propagate data and cancel operations (HTTP requests, database operations, etc.) across goroutines.
For example, using a context you can pass user_id data gathered from a JWT token to child goroutines. 

You can also create a timeout context and set its timeout value to 3 seconds, then use that context in a function that makes a request to a third-party service which shouldn't be waited on for more than 3 seconds. You can inspect the `example1.go` file for further examples.

# Context Cancellation

In Go, when a channel is closed, all channel listeners receive the zero value of that channel's type. This behavior is the foundation of how the Done channel works in contexts. 

The Done channel in a context is a `chan struct{}` type. When a context is cancelled (through manual cancellation, timeout, or deadline), its Done channel is closed. This closure makes all receivers get the zero value (empty `struct{}`), signaling that the context is no longer valid.

