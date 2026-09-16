# Go Code Organization

## Demo Layout

Store independent exercises under:

```text
demos/week-XX/day-XX-topic/
├── go.mod                 # only when the demo needs its own module
├── main.go                # for runnable command demos
├── *_test.go              # tests next to the code under test
└── README.md              # goal, run command, and learned points
```

Keep the first version of a demo small. A single package is preferred until splitting packages demonstrates a real boundary.

## Data and Behavior

- Use `struct` to model state and methods to express behavior owned by that state.
- Prefer composition and small interfaces over inheritance-like abstractions.
- Introduce an interface when there are multiple implementations, a useful test double, or a clear dependency boundary.
- Use `NewType` functions only when construction needs validation or setup; do not create constructors mechanically.
- Keep exported identifiers minimal. Export a name only when another package or the exercise's public API needs it.

## Naming

- Use short, descriptive lower-case package names.
- Use mixedCaps for identifiers; initialisms follow common Go spelling such as `ID`, `URL`, and `HTTP`.
- Name errors with an `Err` prefix when they are package-level sentinel errors.
- Keep one concept per demo and name the directory after the behavior being practiced.

## Avoid

- Do not add a service layer, repository layer, or interface only to imitate a framework.
- Do not hide simple data flow behind unnecessary helper functions.
- Do not copy a production-scale directory structure into a beginner exercise.
