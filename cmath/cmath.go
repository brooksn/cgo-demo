package cmath

/*
#include <calc.h>
*/
import "C"

func Add(a int, b int) int {
	c := C.add(C.int(a), C.int(b))
	return int(c)
}
