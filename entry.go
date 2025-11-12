package optionsgo

import (
	"github.com/yaadata/optionsgo/core"
	"github.com/yaadata/optionsgo/internal"
)

// None creates an Option that contains no value.
//
// Use None when you want to represent the absence of a value//
// Example:
//
//	// Function that may not return a value
//	func findUser(id int) Option[User] {
//	    if id < 0 {
//	        return None[User]()
//	    }
//	    // ... find user logic
//	    return Some(user)
//	}
//
//	result := findUser(-1)
//	if result.IsNone() {
//	    fmt.Println("User not found")
//	}
func None[T any]() core.Option[T] {
	return internal.None[T]()
}

// Some creates an Option that contains the provided value.
//
// Use Some when you have a valid value to wrap in an Option type.
// Example:
//
//	// Function that returns a value when found
//	func findUser(id int) Option[User] {
//	    if id < 0 {
//	        return None[User]()
//	    }
//	    // ... find user logic
//	    return Some(user)
//	}
//
//	result := findUser(123)
//	if result.IsSome() {
//	    user := result.Unwrap()
//	    fmt.Printf("Found user: %v\n", user)
//	}
func Some[T any](value T) core.Option[T] {
	return internal.Some(value)
}

// ResultFromReturn converts Go's standard (value, error) return pattern into a Result.
// If err is not nil, returns Err. Otherwise, returns Ok with the value.
//
// This function is particularly useful for wrapping existing Go functions that follow
// the conventional (T, error) return pattern.
//
// Example:
//
//	type User struct { name string }
//
//	func getUser() (*User, error) {
//	    return &User{name: "Alice"}, nil
//	}
//
//	result := ResultFromReturn(getUser())
//	if result.IsOk() {
//	    user := result.Unwrap() // &User{name: "Alice"}
//	}
//
//	func failedOperation() (*User, error) {
//	    return nil, errors.New("not found")
//	}
//
//	result := ResultFromReturn(failedOperation())
//	if result.IsError() {
//	    err := result.UnwrapErr() // errors.New("not found")
//	}
//
//	// Even with nil value and nil error, returns Ok
//	func nilReturn() (*User, error) {
//	    return nil, nil
//	}
//
//	result := ResultFromReturn(nilReturn())
//	result.IsOk() // true
//	result.Unwrap() // nil
func ResultFromReturn[T any](value T, err error) core.Result[T] {
	return internal.ResultFromReturn(value, err)
}

// Err creates a Result containing an error.
//
// Example:
//
//	result := Err[string](errors.New("something went wrong"))
//	result.IsError() // true
//	result.UnwrapErr() // errors.New("something went wrong")
//	result.UnwrapOr("default") // "default"
func Err[T any](err error) core.Result[T] {
	return internal.Err[T](err)
}

// Ok creates a Result containing a successful value.
//
// Example:
//
//	result := Ok("success")
//	result.IsOk() // true
//	result.Unwrap() // "success"
//	result.UnwrapOr("default") // "success"
func Ok[T any](value T) core.Result[T] {
	return internal.Ok(value)
}
