package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	id         string
	operations []float64
}

func main() {
	cust1 := customer{id: "x"}
	cust2 := customer{id: "x", operations: []float64{}}
	fmt.Println(reflect.DeepEqual(cust1, cust2))
}
