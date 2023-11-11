package main

import "fmt"

func main() {
	var x int = 10 //10
	var y = 10     // int 10
	var z int      // 0

	fmt.Printf("(x,y,z) = (%d,%d,%d)\n", x, y, z)

	var a, b int = 10, 20
	var c, d int
	var e, f = 10, "hello"

	fmt.Printf("(a,b,c,d,e,f) = (%d,%d,%d,%d,%d,%s)\n", a, b, c, d, e, f)

	var (
		h    int
		i        = 20
		j    int = 30
		k, l     = 40, "goodbye"
		m, n string
	)

	fmt.Printf("(h,i,j,k,l,m,n) = (%d,%d,%d,%d,%s,%s,%s)\n", h, i, j, k, l, m, n)

	o, p := 30, "salut"
	q := 23
	o, r := 84, "hallo"

	fmt.Printf("(o,p,q,r) = (%d,%s,%d,%s)\n", o, p, q, r)
}
