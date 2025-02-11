# Go Context Examples

This repository demonstrates practical examples of using contexts in Go for managing cancellation signals, timeouts, and data propagation across goroutines.

## Examples Overview

1. **HTTP Request Timeout** (example1.go)
   - Demonstrates making concurrent HTTP requests with timeout controls
   - Shows how to properly cancel requests that exceed the time limit

2. **Task Timeout** (example2.go)
   - Illustrates basic task timeout implementation
   - Uses select statements for handling timeouts gracefully

3. **Context Values** (example3.go)
   - Shows how to pass values through context
   - Demonstrates sharing data across goroutines safely

4. **Context Deadline** (example4.go)
   - Implements deadline-based cancellation
   - Examples of handling context cancellation signals

5. **HTTP Client Timeout** (example5.go)
   - Advanced HTTP client implementation with context
   - Shows proper request cancellation handling

6. **Database Operations** (example6.go)
   - Demonstrates using context with SQL operations
   - Includes CRUD operations with timeout controls

7. **Custom Timeout Cause** (example7.go)
   - Shows how to use WithTimeoutCause
   - Implements custom error messages for timeouts

## Running the Examples

Use the provided Makefile to run specific examples:

```bash
make run 1  # Runs example1.go
make run 2  # Runs example2.go
# ... and so on
```

## Key Concepts

- Contexts provide a way to carry deadlines, cancellation signals, and request-scoped values across API boundaries
- They're essential for controlling timeouts in distributed systems
- Help prevent resource leaks by properly cancelling operations

## Requirements

- Go 1.20 or higher
- SQLite (for database examples)
```

This README.md is more structured, concise, and provides a clear overview of what each example demonstrates while maintaining brevity. It includes the essential information needed to understand and run the examples.