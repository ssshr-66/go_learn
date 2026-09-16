# Go Testing and Tooling

## Required Commands

Use these commands as the normal feedback loop:

```bash
gofmt -w .
go test ./...
go test -race ./...       # concurrency demos
go vet ./...              # when the package layout supports it
```

Run the narrowest useful command while iterating, then run the full command before completing a demo.

## Tests

- Test behavior through the public API rather than private implementation details.
- Cover the normal path, invalid input, missing data, repeated calls, and lifecycle boundaries when relevant.
- Prefer table-driven tests once several cases share the same assertion shape.
- Make concurrency tests synchronize on observable events instead of relying only on sleeps.
- Give tests names that explain the behavior being protected.

## Debugging

When a test fails:

1. Read the complete error and identify the first failing assertion.
2. Reproduce it with the smallest test or command.
3. State the expected and actual behavior before changing code.
4. Make the smallest fix and add or update a regression test.

## AI-Assisted Work

The learner's first implementation is the primary assessment. AI review is allowed after that attempt. Generated code is not considered learned until the learner can explain it, modify it, and reproduce the relevant tests independently.
