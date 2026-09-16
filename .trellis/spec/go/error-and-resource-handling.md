# Go Error and Resource Handling

## Errors

- Return expected failures explicitly as `error`; do not use `panic` for invalid user input or ordinary operational failures.
- Check errors at the point where the program can make a decision about them.
- Add context while preserving the cause with wrapping such as `fmt.Errorf("load task: %w", err)`.
- Keep error messages lower-case and avoid duplicating context at every layer.
- Use a sentinel or custom error only when callers need to distinguish a condition programmatically.

## Resources

- Register cleanup immediately after successful acquisition with `defer`.
- Make ownership clear: the function that acquires a resource normally owns closing it unless the API explicitly transfers ownership.
- Treat cancellation, timeout, and shutdown as normal lifecycle paths that need tests.

## Learning Checkpoints

For each demo, explain:

1. Which failures are expected and how they are returned.
2. Which function owns each resource.
3. Whether callers can inspect or wrap the returned error.

## Avoid

- Do not ignore an error with `_` unless the reason is explicit and harmless.
- Do not log an error and then return the same error without adding useful context.
- Do not use `panic` as a substitute for designing an error path.
