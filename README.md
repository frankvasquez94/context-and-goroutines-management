# Context and goroutines management

## Problem
Run subprocesses in goroutines such as audit logs in database, cleanup operations, etc. without affecting the service response time.

## Challenge
- Avoid to cancel goroutines because of context cancellation.
- Preserve context data in goroutines.

## Solutions

- Use an indenpendent contex.Background for goroutines. This solution loses context data for a request.
- Implement `Detached context pattern`. This solution creates a custom context from a parent while preserving the parent data.
