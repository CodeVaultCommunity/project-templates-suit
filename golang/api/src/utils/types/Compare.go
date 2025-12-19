// Package utilstypes Contains types for comparation.
package utilstypes

// CompareFunc Recives two arguments, a and b, that such have the same type T
// and compare them; returns 1 if a is greater than b, -1 if b is greater than a
// and 0 if a and b ar equal
type CompareFunc[T any] = func(a T, b T) int

// Comparator is a interface to create objects that know how to compare tow objects
type Comparator[T any] interface {
	/*
		Compare same as CompareFunction:

			CompareFunc Recives two arguments, a and b, that such have the same type T
			and compare them; returns 1 if a is greater than b, -1 if b is greater than a
			and 0 if a and b ar equal
	*/
	Compare(a T, b T) int
}
