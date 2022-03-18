package main

import "fmt"

type customer struct {
	id         string
	operations []float64
}

func main() {
	cust1 := customer{id: "x", operations: []float64{1.}}
	cust2 := customer{id: "x", operations: []float64{1.}}
	fmt.Println(cust1 == cust2)
}
