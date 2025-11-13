package core

// OptionChain provides a chainable interface for transforming Option values.
// It allows multiple operations to be composed fluently before converting back to an Option.
type OptionChain[T any] interface {
	// Option converts the OptionChain back to an Option[T].
	// This is typically called at the end of a chain to retrieve the final Option value.
	//
	// Example:
	//  result := Some(15).
	//      Filter(func(val int) bool { return val > 10 }).
	//      Map(func(value int) any { return fmt.Sprintf("Value=%d", value) }).
	//      Option()  // Converts back to Option[any]
	//  result.Unwrap() // "Value=15"
	Option() Option[T]

	// AndThen chains another option-returning operation.
	// If the current chain represents Some, applies fn to the value and returns a new OptionChain with the result.
	// If the current chain represents None, returns an OptionChain representing None without calling fn.
	//
	// This is useful for chaining operations that might fail and return None.
	//
	// Example:
	//  result := Some(3).
	//      Map(func(v int) any { return v * 2 }).
	//      AndThen(func(v any) Option[any] {
	//          if num, ok := v.(int); ok && num > 5 {
	//              return Some(num)
	//          }
	//          return None[any]()
	//      }).
	//      Option()
	AndThen(fn func(T) Option[any]) OptionChain[any]

	// Map transforms the value in the chain by applying a function.
	// If the current chain represents Some, applies fn to the value and returns a new OptionChain with the transformed value.
	// If the current chain represents None, returns an OptionChain representing None without calling fn.
	//
	// This enables fluent transformation of values while maintaining the Option context.
	//
	// Example:
	//  result := Some(15).
	//      Filter(func(val int) bool { return val > 10 }).
	//      Map(func(value int) any { return fmt.Sprintf("Value=%d", value) }).
	//      Option()
	//  result.Unwrap() // "Value=15"
	//
	//  none := None[int]().
	//      Map(func(value int) any { return fmt.Sprintf("Value=%d", value) }).
	//      Option()
	//  none.IsNone() // true
	Map(fn func(T) any) OptionChain[any]

	// MapOr transforms the value or returns a default value, terminating the chain.
	// If the current chain represents Some, applies fn to the value and returns the transformed result.
	// If the current chain represents None, returns the provided default value 'or' without calling fn.
	//
	// This method terminates the chain and returns the final value directly (not an Option).
	//
	// Example:
	//  result := None[string]().
	//      Map(func(a string) any { return len(a) }).
	//      MapOr(func(value any) any {
	//          if val, ok := value.(int); ok {
	//              return fmt.Sprintf("VALUE=%d", val)
	//          }
	//          return ""
	//      }, "OTHER")
	//  result // "OTHER"
	MapOr(fn func(T) any, or any) any

	// MapOrElse transforms the value or computes a default value, terminating the chain.
	// If the current chain represents Some, applies fn to the value and returns the transformed result.
	// If the current chain represents None, calls orElse to compute a default value and returns that without calling fn.
	//
	// This method terminates the chain and returns the final value directly (not an Option).
	// Use this when computing the default value is expensive and should only happen when needed.
	//
	// Example:
	//  result := None[string]().
	//      Map(func(a string) any { return len(a) }).
	//      MapOrElse(func(value any) any {
	//          if val, ok := value.(int); ok {
	//              return fmt.Sprintf("VALUE=%d", val)
	//          }
	//          return ""
	//      }, func() any {
	//          return "OTHER"
	//      })
	//  result // "OTHER"
	MapOrElse(fn func(T) any, orElse func() any) any
}
