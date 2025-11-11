package optionsgo

// Result represents the outcome of an operation that can either succeed with a value
// or fail with an error. This is a Go implementation of Rust's std::result::Result type.
//
// A Result is either:
//   - Ok: contains a successful value of type T
//   - Err: contains an error
type Result[T any] interface {
	// Ok returns Some(value) if the result is Ok, otherwise returns None.
	//
	// Example:
	//
	//  result := Ok("value")
	//  opt := result.Ok() // Some("value")
	//
	//  result := Err[string](errors.New("error"))
	//  opt := result.Ok() // None
	Ok() Option[T]

	// Err returns Some(error) if the result is Err, otherwise returns None.
	//
	// Example:
	//
	//  result := Err[string](errors.New("error_message"))
	//  opt := result.Err() // Some(error)
	//
	//  result := Ok("value")
	//  opt := result.Err() // None
	Err() Option[error]

	// IsOk returns true if the result is Ok.
	//
	// Example:
	//
	//  result := Ok("value")
	//  result.IsOk() // true
	//
	//  result := Err[string](errors.New("error"))
	//  result.IsOk() // false
	IsOk() bool

	// IsOkAnd returns true if the result is Ok and the predicate returns true for the value.
	//
	// Example:
	//
	//  result := Ok("value")
	//  result.IsOkAnd(func(s string) bool { return len(s) > 0 }) // true
	//  result.IsOkAnd(func(s string) bool { return len(s) == 0 }) // false
	//
	//  result := Err[string](errors.New("error"))
	//  result.IsOkAnd(func(s string) bool { return true }) // false
	IsOkAnd(pred Predicate[T]) bool

	// IsError returns true if the result is Err.
	//
	// Example:
	//
	//  result := Err[string](errors.New("error"))
	//  result.IsError() // true
	//
	//  result := Ok("value")
	//  result.IsError() // false
	IsError() bool

	// IsErrorAnd returns true if the result is Err and the predicate returns true for the error.
	//
	// Example:
	//
	//  result := Err[string](errors.New("error_message"))
	//  result.IsErrorAnd(func(e error) bool { return e.Error() == "error_message" }) // true
	//  result.IsErrorAnd(func(e error) bool { return e.Error() == "other" }) // false
	//
	//  result := Ok("value")
	//  result.IsErrorAnd(func(e error) bool { return true }) // false
	IsErrorAnd(pred Predicate[error]) bool

	// Unwrap returns the contained Ok value.
	// Panics if the result is Err.
	//
	// Example:
	//
	//  result := Ok("EXPECTED")
	//  value := result.Unwrap() // "EXPECTED"
	//
	//  result := Err[string](errors.New("error"))
	//  value := result.Unwrap() // panics!
	Unwrap() T

	// UnwrapErr returns the contained Err value.
	// Panics if the result is Ok.
	//
	// Example:
	//
	//  err := errors.New("error_message")
	//  result := Err[string](err)
	//  e := result.UnwrapErr() // err
	//
	//  result := Ok("value")
	//  e := result.UnwrapErr() // panics!
	UnwrapErr() error

	// UnwrapOr returns the contained Ok value or the provided default value.
	//
	// Example:
	//
	//  result := Ok("EXPECTED")
	//  value := result.UnwrapOr("default") // "EXPECTED"
	//
	//  result := Err[string](errors.New("error"))
	//  value := result.UnwrapOr("default") // "default"
	UnwrapOr(value T) T

	// UnwrapOrElse returns the contained Ok value or computes it from the provided function.
	//
	// Example:
	//
	//  result := Ok("EXPECTED")
	//  value := result.UnwrapOrElse(func() string { return "default" }) // "EXPECTED"
	//
	//  result := Err[string](errors.New("error"))
	//  value := result.UnwrapOrElse(func() string { return "default" }) // "default"
	UnwrapOrElse(fn func() T) T

	// UnwrapOrDefault returns the contained Ok value or the zero value of type T.
	//
	// Example:
	//  result := Ok("EXPECTED")
	//  value := result.UnwrapOrDefault() // "EXPECTED"
	//
	//  result := Err[string](errors.New("error"))
	//  value := result.UnwrapOrDefault() // ""
	UnwrapOrDefault() T
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
func ResultFromReturn[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return &result[T]{
		value: &value,
		err:   nil,
	}
}

// Err creates a Result containing an error.
//
// Example:
//
//	result := Err[string](errors.New("something went wrong"))
//	result.IsError() // true
//	result.UnwrapErr() // errors.New("something went wrong")
//	result.UnwrapOr("default") // "default"
func Err[T any](err error) Result[T] {
	return &result[T]{
		value: nil,
		err:   err,
	}
}

// Ok creates a Result containing a successful value.
//
// Example:
//
//	result := Ok("success")
//	result.IsOk() // true
//	result.Unwrap() // "success"
//	result.UnwrapOr("default") // "success"
func Ok[T any](value T) Result[T] {
	return &result[T]{
		value: &value,
		err:   nil,
	}
}
