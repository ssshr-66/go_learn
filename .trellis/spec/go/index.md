# Go Engineering Learning Guidelines

## Scope

These guidelines apply to all Go learning demos and exercises in this repository. They describe the working conventions for a progressive learning project, not the conventions of an external production codebase.

## Pre-Development Checklist

Before writing a demo:

1. Read the active weekly task and identify the single concept being practiced.
2. Write down the input, output, state, and failure cases in plain language.
3. Start with the smallest runnable implementation; defer abstractions until the exercise needs them.
4. Attempt the current checkpoint before requesting a complete implementation from AI.
5. Decide how the result will be run and tested before adding extra code.

## Quality Check

Every completed demo should:

- be formatted with `gofmt`;
- run with `go run .` or have a clear test-only entry point;
- have tests for the normal path and important failure or boundary cases;
- pass `go test ./...`;
- include a short explanation of the design and a note about any AI assistance.

Concurrency demos additionally run `go test -race ./...` when the code is ready for it.

## Guidelines Index

| Guide | Purpose |
|---|---|
| [Code Organization](./code-organization.md) | Demo layout, packages, structs, methods, interfaces, and composition |
| [Error and Resource Handling](./error-and-resource-handling.md) | Explicit errors, wrapping, `defer`, and cleanup |
| [Testing and Tooling](./testing-and-tooling.md) | Formatting, tests, debugging, and verification commands |
| [Concurrency](./concurrency.md) | Goroutines, channels, synchronization, cancellation, and shutdown |

## Learning Boundary

Do not introduce Spring-style frameworks, deep runtime internals, reflection, advanced generics, or large external dependencies until a small standard-library demo has made the underlying problem concrete.
