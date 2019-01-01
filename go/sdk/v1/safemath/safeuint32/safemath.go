package safeuint32

import "math"

func Add(x uint32, y uint32) uint32 {
	if y > math.MaxUint32-x {
		panic("integer overflow on add")
	}
	return x + y
}

func Sub(x uint32, y uint32) uint32 {
	if x < y {
		panic("integer overflow on sub")
	}
	return x - y
}

func Mul(x uint32, y uint32) uint32 {
	if x == 0 || y == 0 {
		return 0
	}
	if y > math.MaxUint32/x {
		panic("integer overflow on mul")
	}
	return x * y
}

func Div(x uint32, y uint32) uint32 {
	if y == 0 {
		panic("division by zero")
	}
	return x / y
}

func Mod(x uint32, y uint32) uint32 {
	if y == 0 {
		panic("division by zero")
	}
	return x % y
}
