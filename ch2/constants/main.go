package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(idKey)
	fmt.Println(nameKey)

	const w = 10

	var a int = w
	var b float64 = w
	var c byte = w
	fmt.Printf("(w,a,b,c) = (%d,%d,%f,%d)", w, a, b, c)
}
