package day1

import "fmt"

func VarDef() {

  // The most verbose way to assign a variable in Go
  var x int = 30
  fmt.Println("The value of x is", x)

  // we can leave out the type definition if we want.
  // go compiler will see there is a integer literal and assign a type to variable y
  var y = 20
  fmt.Println("The value of y is", y)

  // If we want to define a variable and set it to it's zeroth value we can do this
  var z int // here z would be set to 0 
  fmt.Println("The value of z is", z)

  // We can also define multiple variables at the same time
  var a, b int = 10, 20
  fmt.Println("Multiple variables defined", a, b)

  // We can default them to their zeroth values
  var c, d int
  fmt.Println("Multiple variables defined with zeroth values", c, d)

  // We can also define different types
  var e,f = 10, "Hello"
  fmt.Println("Multiple variables defined, with their own types", e, f)

  // If we want to define a lot of multiple vaiables, we can wrap them in a declaration list
  // Similar to what we do when we import multiple modules
  var (
    first, last string
    age int
    height float64
    weight float64
  )
  fmt.Println("Variables defined in a declaration list:", first, last, age, height, weight)

  // Go also supports the opposite of declartion list, a short declaration. Where you can
  // do away with var keyword itself.
  // The only restriction is that this kind of definition can only be inside a function
  // Notice the operator := instead of =
  sh := 10
  fmt.Println("Variable defined using short declaration", sh)

  /*
  The := operator can do one trick that you cannot do with var: it allows you to assign values to existing variables, too. As long as there is one new variable on the lefthand side of the :=, then any of the other variables can already exist: 
  */
  sh, ph := 20, "Name"
  fmt.Println("Declaring new and assigning values to old variables using short declartion", sh, ph)


}