package main

import (
	"fmt"

	"github.com/brooksn/cgo-demo/cmath"
)

func main() {
	a := 1
	b := 2
	c := cmath.Add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, c)
}
