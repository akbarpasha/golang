package main

import "fmt"

type customer struct {
	id         string
	operations []float64
}

func main() {
	var cust1 interface{} = customer{id: "x", operations: []float64{1.}}
	var cust2 interface{} = customer{id: "x", operations: []float64{1.}}
	fmt.Println(cust1 == cust2)
}
