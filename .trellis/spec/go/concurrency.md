# Go Concurrency Guidelines

These rules apply after the language and testing foundations are in place.

## Core Rules

- State the ownership of every goroutine, channel, and shared variable before coding.
- Prefer a fixed number of workers or another bounded design when work can arrive faster than it is processed.
- The sender or component that owns a channel's lifecycle is responsible for closing it; receivers must not close a channel they do not own.
- Use `sync.WaitGroup` to wait for goroutines, and define exactly where `Add`, `Done`, and `Wait` occur.
- Use `context.Context` for cancellation and deadlines that cross API boundaries.
- Run the race detector for concurrent demos and test repeated shutdown, cancellation, and empty-work cases.

## Learning Order

1. Goroutine and channel basics
2. Buffered versus unbuffered channels and `select`
3. `Mutex`, `RWMutex`, `WaitGroup`, and atomic operations
4. `context` cancellation and timeouts
5. Worker pools, bounded queues, retries, and graceful shutdown

## Avoid

- Do not use sleeps as the primary synchronization mechanism.
- Do not start a goroutine without a clear exit condition.
- Do not mix channel coordination and shared mutable state without explaining why both are needed.
- Do not optimize for throughput before measuring correctness, leaks, and race behavior.
