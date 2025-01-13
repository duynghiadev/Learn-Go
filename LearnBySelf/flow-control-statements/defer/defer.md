# Differences Between `defer` in Go and `async/await` in JavaScript

While `defer` in Go and `async/await` in JavaScript might seem similar at first glance because they both deal with execution flow, they are quite different in terms of functionality and use cases.

## Key Differences

### Execution Timing

- **`defer` in Go**:
  The deferred function call is executed after the surrounding function finishes executing, but before the function returns. It is a way to schedule cleanup or finalization tasks (like closing files or unlocking resources) that should happen at the end of a function’s execution.

- **`async/await` in JavaScript**:
  `async` is used to define a function that returns a promise, and `await` is used to pause the execution of the function until the promise is resolved or rejected. This is used for asynchronous operations, such as I/O tasks, network requests, or other non-blocking operations.

### Purpose

- **`defer`**:
  Primarily used for cleanup tasks after a function completes (like closing files, releasing locks, or releasing resources).

- **`async/await`**:
  Used to handle asynchronous operations, ensuring that certain tasks (like HTTP requests or database queries) complete before proceeding with the rest of the code.

### Execution Flow

- **`defer`**:
  The deferred function runs after the function’s main execution completes, but before the function exits. Multiple `defer` calls are executed in LIFO (Last In, First Out) order.

- **`async/await`**:
  `await` pauses the function execution until the promise resolves, allowing other code to run in the meantime (non-blocking). It’s used to write asynchronous code that behaves like synchronous code.
