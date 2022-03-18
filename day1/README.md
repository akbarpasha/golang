# Primitive Types and Declarations
### Zero Values
- Go, like most modern languages, assigns a default zero value to any variable that is declared but not assigned a value.
- The zero value of a numeric type is 0, for booleans it is false, and for strings it is the empty string.
- The zero value of a pointer type is nil.
- The zero value of a struct is the zero value of each of its fields.
- The zero value of an interface value is nil.
- The zero value of a map is nil.
- The zero value of an array is nil.
- The zero value of a channel is nil.
- The zero value of a function is nil.
- The zero value of a slice is nil.
### Literals
- A literal in Go refers to writing out a number, character, or string. There are four common kinds of literals that you’ll find in Go programs.
     - Integer literals are sequence of numbers.
     - Floating-point literals are sequence of numbers with a decimal point.
     - Rune literals are a single character surrounded by single quotes.
          - In Go single quotes and double quotes are not interchangeable.
     - String literals are a sequence of characters surrounded by double quotes.   
     - If you need to include backslashes, double quotes, or newlines in your string, use a raw string literal. These are delimited with backquotes (`) and can contain any literal character except a backquote.
-  Literals in Go are untyped; they can interact with any variable that’s compatible with the literal.

### Booleans
- The bool type represents Boolean variables. Variables of `bool` type can have one of two values: `true` or `false`. The zero value for a `bool` is `false`.

### Numeric Types
- Go has 12 numeric types:
     - int8, int16, int32, int64
     - uint8, uint16, uint32, uint64
     - float32, float64
     - complex64, complex128
- A `byte` is an alias for `uint8`.
- You can assign, compare or perform mathematical operations between `byte` and `uint8` values.
- `int` is an alias for `int64` or `int32` depending on type of CPU architecture.
- `uint` is an alias for `uint64` or `uint32` depending on type of CPU architecture.
- `rune` is an alias for `int32` and represents a Unicode code point.
- `uintptr` is an alias for `uint` and represents an unsigned integer value.

### String Types
- A `string` is a sequence of Unicode characters.
- A string literal is a sequence of characters surrounded by double quotes.
- String are immutable.

### var declaration

- `var` is the most verbose way to declare a variable.
     `var x int = 100`
     
- We can declare multiple variables at once.
     `var x, y, z int = 1, 2, 3`
     
- Or of different types.
     `var x, y = 1, "hello"`
     
- For multiple variables, we can wrap them in declaration lists.
     ```go
     var (
       x int
       y int = 10
       z string
     )
     
- We can also use short declaration syntax.
     `x := 100`. Same as `var x = 100`
     
- The `:=` operator can do one trick that you cannot do with `var`,  it allows you to assign values to existing variables, too. As long as there is one new variable on the lefthand side of the `:=`, then any of the other variables can already exist:

- ```go
     x := 10
     x, y := 30, "hello"
     ```

- There is one limitation on `:=`. If you are declaring a variable at package level, you must use `var` because `:=` is not legal outside of functions.

  

### const

- Constants in Go are a way to give names to literals. They can only hold values that the compiler can figure out at compile time.

- They can be assigned:

  - Numeric literals
  - `true` and `false`
  - Strings
  - Runes
  - The built-in functions `complex`, `real`, `imag`, `len`, and `cap`
  - Expressions that consist of operators and the preceding values

- Go doesn’t provide a way to specify that a value calculated at runtime is immutable.

- There are no immutable arrays, slices, maps, or structs, and there’s no way to declare that a field in a struct is immutable.

- Constants in Go are a way to give names to literals. There is *no* way in Go to declare that a variable is immutable.

- Constants can be typed or untyped. 

  - An untyped constant works exactly like a literal; it has no type of its own, but does have a default type that is used when no other type can be inferred.

  - ```go
    const x = 10
    var y int = x
    var z float64 = x
    var d byte = x
    ```

  -  A typed constant can only be directly assigned to a variable of that type.

  - ```go
    const typedX int = 10
    ```

### Unused variables

- Another Go requirement is that *every declared local variable must be read*. It is a *compile-time error* to declare a local variable and to not read its value.
- As long as a variable is read once, the compiler won’t complain, even if there are writes to the variable that are never read.
- The Go compiler won’t stop you from creating unread package-level variables.
- Perhaps surprisingly, the Go compiler allows you to create unread constants with `const`. This is because constants in Go are calculated at compile time and cannot have any side effects. This makes them easy to eliminate: if a constant isn’t used, it is simply not included in the compiled binary.



Notes from the amazing book - [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/) by [Jon Bodner](https://medium.com/@jon_43067)

