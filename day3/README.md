# Comparing Values

Comparing values in our programs is very common. There are few gotcha's in Go to watchout for when comparing values.

- go uses `==` and `!=`operators for comparison.

### Comparing a struct

```go
type customer struct { 
  id string 
}

func main() { 
  cust1 := customer{id: "x"} 
  cust2 := customer{id: "x"} 
  fmt.Println(cust1 == cust2) 
}
```

- The above comparison - gives out a `true` value.

- If we add in a slice field:

  ```go
  type customer struct { 
    id string 
    operations []float64 
  }
  
  func main() { 
    cust1 := customer{id: "x", operations: []float64{1.}} 
    cust2 := customer{id: "x", operations: []float64{1.}} 
    fmt.Println(cust1 == cust2) 
  }
  ```

- The code above won't even compile. 

- Go, operators `==` and `!=` can be used to compare types that are *comparable*

- Comparable types are:

  - Booleans:  `==` and `!=` to compare whether two booleans are equals.
  - Numerics: `==` and `!=` to compare whether two numerics are equals. These operations compile if both variables have the same type or converted to be of the same type.
  - Strings: `==` and `!=` to compare whether two strings are equals. We can also use , `>=`, `<` and `>` operators to compare two strings based on their lexical order.
  - Pointers: `==` and `!=` to compare whether two pointers point to the same value in memory or if both are `nil`
  - Channels: `==` and `!=` to compare whether two channels were created by the same call to `make` or if both are `nil`.
  - Arrays and Structs can be used with `==` and `!=` operators if they contain purely comparable types.

- We cannot use maps and slices with `==` and `!=`  operators

- There are some issues with comparison operators with `interface{}`

- What if we first cast them into `interface{}` instead of comparing two customers struct?

  ```go
  type customer struct { 
    id string 
    operations []float64 
  }
  
  func main() {
  	var cust1 interface{} = customer{id: "x", operations: []float64{1.}} 
  	var cust2 interface{} = customer{id: "x", operations: []float64{1.}} 
  	fmt.Println(cust1 == cust2)
  }
  ```

- The above code compiles but raises a panic during run time

  `panic: runtime error: comparing uncomparable type main.customer`

- To compare maps and slices, we can use reflection and use `reflect.DeepEqual`

  ```Go
  cust1 := customer{id: "x", operations: []float64{1.}} 
  cust2 := customer{id: "x", operations: []float64{1.}} 
  fmt.Println(reflect.DeepEqual(cust1, cust2)) // returns true
  ```

- There is one caveat as how `reflect.DeepEqual` treats `nils` and empty slice.

  ```go
  var cust1 interface{} = customer{id: "x"} // with nil slice
  var cust2 interface{} = customer{id: "x", operations: []float64{}} // with empty slice
  fmt.Printf("%t\n", reflect.DeepEqual(cust1, cust2)) // returns false
  ```

- `reflect.DeepEqual` considers empty and nil collections different

- Also `reflect.DeepEqual` is slower and has a performance penalty.



[^Notes from the amazing book - [100 Go mistakes and how to avoid them](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them) by [Teiva Harsanyi](https://teivah.github.io/)]:

