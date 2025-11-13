# optionsgo

A Go implementation of Rust's `Option<T>` and `Result<T, E>` types.

## Overview

This library brings Rust's ergonomic error handling patterns to Go through two
core types:

- **`Option[T]`** - Represents an optional value that may or may not exist
- **`Result[T]`** - Represents the outcome of an operation that can either
  succeed or fail

## Dependencies

The following tools are used in this project

- [mise](https://github.com/jdx/mise)

  Version management of Go for local development.

- [just](https://github.com/casey/just)

  Command runner. Alternative to Make.

## Installation

```bash
go get github.com/yaadata/optionsgo
```

## Quick Start

```go
import (
    "errors"
    "fmt"

    "github.com/yaadata/optionsgo"
)

// Using Option
func findUser(id int) optionsgo.Option[string] {
    if id < 0 {
        return optionsgo.None[string]()
    }
    return optionsgo.Some("Alice")
}

result := findUser(123)
if result.IsSome() {
    fmt.Println(result.Unwrap()) // "Alice"
}

// Using Result
func divideNumbers(a, b int) optionsgo.Result[int] {
    if b == 0 {
        return optionsgo.Err[int](errors.New("division by zero"))
    }
    return optionsgo.Ok(a / b)
}

result := divideNumbers(10, 2)
value := result.UnwrapOr(0) // 5
```

## API Reference

### Option[T] Interface

The interface methods can be found in [here](./core/option.go)

View the interface definition via html by running `just doc`

#### Constructor Functions

```go
Some[T](val T) Option[T]     // Creates an Option containing a value
None[T]() Option[T]           // Creates an Option with no value
```

### Result[T] Interface

The interface methods can be found in [here](./core/result.go)

View the interface definition via html by running `just doc`

#### Constructor Functions

```go
Ok[T](value T) Result[T]                           // Creates a Result containing a success value
Err[T](err error) Result[T]                        // Creates a Result containing an error
ResultFromReturn[T](value T, err error) Result[T]  // Converts Go's (T, error) pattern to Result
```

## Usage Examples

### Working with Option[T]

#### Creating and Checking Options

```go
// Create Some option
opt := optionsgo.Some("hello")
opt.IsSome()  // true
opt.IsNone()  // false

// Create None option
empty := optionsgo.None[string]()
empty.IsSome()  // false
empty.IsNone()  // true
```

#### Conditional Checks with Predicates

```go
opt := optionsgo.Some("SOME")

// Check if Some and predicate passes
opt.IsSomeAnd(func(s string) bool {
    return len(s) == 4
}) // true

opt.IsSomeAnd(func(s string) bool {
    return len(s) == 3
}) // false

// Check if None or predicate passes
opt.IsNoneOr(func(s string) bool {
    return len(s) == 4
}) // true (predicate passes)

optionsgo.None[string]().IsNoneOr(func(s string) bool {
    return false
}) // true (is None)
```

#### More Examples ?

See the respective test files

## Commands

### Test

Run the test suite:

```bash
just test
```

Each merge to main is validated on the last 2 major version of Go across
Windows, Mac and Linux.

### Docs

View docs

```bash
just doc
```

### Lint

Run the project linter

```bash
just lint
```

## Contributing

Contributions are welcome! If you'd like to implement any of the missing APIs or
have suggestions for Go-idiomatic alternatives, please open a discussion.

## References

- [Rust std::option::Option](https://doc.rust-lang.org/std/option/enum.Option.html)
- [Rust std::result::Result](https://doc.rust-lang.org/std/result/enum.Result.html)

## License

See LICENSE file for details.
