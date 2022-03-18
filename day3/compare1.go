package main

import "fmt"

type customer struct {
	id string
}

func main() {
	cust1 := customer{id: "x"}
	cust2 := customer{id: "x"}
	fmt.Println(cust1 == cust2)
}
