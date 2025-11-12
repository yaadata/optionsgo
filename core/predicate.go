package core

type Predicate[T any] func(val T) bool
