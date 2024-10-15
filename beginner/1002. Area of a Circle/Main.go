package main

import (
	"fmt"
)

func main() {
	var R float64

	n := 3.14159

	fmt.Scan(&R)

	A := n * R * R

	fmt.Printf("A=%.4f\n", A)
}