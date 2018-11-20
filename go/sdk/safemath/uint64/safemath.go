package uint64

import "math"

func Add(x uint64, y uint64) uint64 {
	if y > math.MaxUint64-x {
		panic("integer overflow on add")
	}
	return x + y
}

func Sub(x uint64, y uint64) uint64 {
	if x < y {
		panic("integer overflow on sub")
	}
	return x - y
}

func Mul(x uint64, y uint64) uint64 {
	if x == 0 || y == 0 {
		return 0
	}
	if y > math.MaxUint64/x {
		panic("integer overflow on mul")
	}
	return x * y
}
