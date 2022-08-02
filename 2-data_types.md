# Data Types

In this post we'll cover some of the data types used in Go.

By way of a comparison with Python, here are some of the the basic data types which you can see are similar:

| Python | Go                             | Examples              |
| ------ | ------------------------------ | --------------------- |
| int    | int, int8, int16, int32, int64 | -128, 0, 42           |
| float  | float32, float64               | -1.12, 0, 3.14159     |
| bool   | bool                           | true, false           |
| str    | string                         | "Hello"               |
| bytes  | byte, []byte                   | 0xff, []byte("Hello") |

We can use the `%T` format specifier for `fmt.Printf` to see the type of the variable. There are many other format specifiers or "verbs" you can use as per [the docs](https://pkg.go.dev/fmt) for the fmt package.

When assigning variables to types, you can be explicit with the variable declaration or use the short form.  These are all equivalent:

```go
i := 0        // 1
var i int     // 2
var i = 0     // 3
var i int = 0 // 4
```

The first(1) short form is preferred typically since it is, well, shorter.  It can only be used inside functions though, not for package level variables.  Also in this case, `int` is implied, so if you need an explicit type or you want to assign the value later, you'd use (2).

Also of note in the second form(2) is that all types in Go have what is called a "zero value" which is the value assigned to variables in the absence of one being specified for them on initialisation.  For integers, this is `0`, for strings, the empty string `""` and for booleans it's `false`.  This is something to watch out for, which we'll probably cover in a future "gotchas" session.

## Numerical Operators

The use of numerical operators is quite common in python, so here is a comparison:


| Operation      | Python | Go  |
| -------------- | :----: | :---: |
| Addition       |   +    | +
| Subtraction    |   -    | -
| Multiplication |   *    | *
| Division       |   /    | /
| Floor Division  |   //   | math.Floor()
| Modulo         |   %    | %
| Power          |   **   | math.Pow()

## Type Conversion

Due to Go being statically typed, it is reasonably common to perform type conversion.  This might seem painful compared to python but at least you know what you're getting if you're in control of it.

As an example, of note in the numerical operators table above is that `math.Floor()` and `math.Pow()` both require the use of `float64` variables, so you can't pass in an `int`.  You would first have to convert an int to a float64 prior to using it, e.g:

```go
math.Pow(float64(x), float64(y))
```

There are a wide range of conversions available and we'll talk some more about those for strings in the next session.

