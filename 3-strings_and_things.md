# Strings and things

This is a dedicated post related specifically to string manipulation since it's something that occurs quite
a lot and will provide an easy reference.

- [Packages](#packages)
  - [fmt package](#fmt-package)
  - [strings package](#strings-package)
  - [strconv package](#strconv-package)
- [Appendices](#appendices)
  - [Resources](#resources)
  - [Errors](#errors)


Strings are enclosed in double quotes only, unlike Python which lets you use single or double quotes. In addition, strings are stored as a slice of bytes (see [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings) for more detail) -- this is important if you're thinking about looping over a string which we cover in [5 - Going Loopy](5-going-loopy.md).

As covered in [data-types](2-data_types.md), strings can be initialised in different ways:

```go
// string initialisation
var s1 string
s1 = "hello"
s2 := "world"
// note that Println inserts a space between each passed variable
fmt.Println(s1, s2)
fmt.Printf("%T, %T\n", s1, s2)
```

Fairly typically, as with most languages, strings are [immutable](https://en.wikipedia.org/wiki/Immutable_object), so whilst you can do:

```go
hello := "Hello, " + "world!"
```

It is not the most efficient way to concatenate strings both in terms of time and space due to having to reallocate memory.  There is almost always a better way.  

This may often be a [strings.Builder](https://pkg.go.dev/strings#Builder):

```go
var hello strings.Builder // remember to import strings
hello.WriteString("Hello, ")
hello.WriteString("world!")
fmt.Println(hello.String())
```

Clearly for smaller strings, this isn't such a problem, but bear it in mind if you start concatenating lots of information to a single string.

You can also combine strings and variables using [fmt.Sprintf](https://pkg.go.dev/fmt#Sprintf)

```go
world := "world"
hello := fmt.Sprintf("Hello, %s!", world)
```

Remember to use the correct [format specifier](https://pkg.go.dev/fmt) for the type, e.g.:

```go
hundred := 100
pi := 3.14159
result := fmt.Sprintf("One hundred: %d and Pi: %.2f", hundred, pi)
fmt.Println(result)
```

Notice that `fmt.Sprintf` results in a string which is assigned to a variable, it doesn't print anything to the console by itself.  Also note that we specified 2 decimal places for pi, so only `3.14` is printed.

Since we were printing to the console we could have combined those statements using [fmt.Printf](https://pkg.go.dev/fmt#Printf):

```go
hundred := 100
pi := 3.14159
fmt.Printf("One hundred: %d and Pi: %.2f\n", hundred, pi)

```

> *Notice now that we need to provide our own newline character*

The use of `f` at the end of the function name indicates that we can use the **f**ormat specifiers.

## Packages

In addition to the `fmt` package, there are various other useful packages when it comes to strings:

* [fmt](https://pkg.go.dev/fmt) - formatted I/O functions
* [strings](https://pkg.go.dev/strings) - simple functions to manipulate UTF-8 encoded strings
* [strconv](https://pkg.go.dev/strconv) - conversions to and from string representations of basic date types

To use these, because they are part of the standard library, we can import them as follows:

```go
import (
    fmt
    strings
    strconv
)
```

### fmt package

You've already seen some examples of the `fmt` package in use, but here are another couple of useful examples:

The use of [fmt.Sscanf](https://pkg.go.dev/fmt#Sscanf) for extracting variables out of already formatted text:

```go
// useful for scanning input text into variables
s := "Bob is 195cm tall and weighs 89kg"
var name string
var height, weight int
fmt.Sscanf(s, "%s is %dcm tall and weighs %dkg", &name, &height, &weight)
fmt.Println(name, height, weight)
```

This syntax uses the addresses of name, height and weight to put the resulting values into the memory allocated for them (we'll cover this more in future), but hopefully you can see how this might be useful when importing data to extract the values you want.

### strings package

The `strings` package is useful for splitting strings, finding the index of a particular substring etc. -- here are some useful examples:

```go
strings.HasPrefix("https://www.cisco.com","http")                 // true
strings.HasSuffix("https://www.cisco.com","COM")                  // false, case sensitive
strings.ToLower("COM")                                            // "com"
strings.Split("www.cisco.com",",")                                // ["www","cisco","com"]
strings.Fields("   ip dhcp excluded-address 10.0.1.1 10.0.1.10 ") // ["ip", "dhcp", "excluded-address", "10.0.1.1", "10.0.1.10"]
strings.Join([]string{"enable", "password", "cisco123"}," ")      // "enable password cisco123"
fmt.Println("ba" + strings.Repeat("na", 2))                       // "banana"
```

> *Note that some of these used or resulted in "slices of strings" which are like lists in python. We'll cover those more in the next post.*

You can see all of the functions available by looking at the [documentation for the strings library](https://pkg.go.dev/strings).

### strconv package

The `strconv` package is useful when you need to convert other object types to string or strings to other object types and you're not just printing them to the console.

To convert strings to values:

```go
b, err := strconv.ParseBool("true")
f, err := strconv.ParseFloat("3.1415", 64)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)
```

> *See the [Appendix on Errors](#errors) for detail on the `err` syntax here*

To convert values to strings:

```go
s := strconv.FormatBool(true)
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
s := strconv.FormatInt(-42, 16)
s := strconv.FormatUint(42, 16)
```

## Appendices

### Resources

* [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings) -- a must read regarding character encoding, UTF-8, Code Points and Runes

### Errors

Unlike Python and many other languages, Go doesn't have the concept of "exceptions", so there is no `try/catch` syntax. 

Instead many functions will return an `error` as a second response parameter, e.g.:

```go
b, err := strconv.ParseBool("true")
fmt.Println(b, err) // true, nil
```

In this case, there is no error since the string value "true" converted to a boolean just fine.  However, if we pass something else:

```go
b, err := strconv.ParseBool("networking")
fmt.Println(b, err) // false, syntax error
```

We will get a syntax error.  Notice that here, `b` will be set to false, which if you remember is the zero value for a boolean, so it's really important to check for errors to ensure you're getting the value you expect.

So you will see this syntax a lot in Go programs. It seems annoying at first, but it means that you have far fewer unhandled errors and you do get used to it:

```go
b, err := strconv.ParseBool("true")
if err != nil {
    // we handle the error at this point
    // e.g. return or exit
    return err
}
// we only get here if the error was nil
fmt.Println("We got what we expected")
```

> *Sidenote: something quite common in Go programs is to not use `else` with an `if` statement. Whilst you can, it makes the code more readable if there are fewer indented blocks, so it's quite common in functions to just return the error you received which then no longer requires the else.*

Finally, you can also ignore (at your peril) errors by using an `_` (underscore) character where an error is provided, since Go will complain if you have an unused error:

```go
b, _ := strconv.ParseBool("true")
```

A separate post on errors might be useful at some point, but you'll see plenty of examples as we go through the series. In the meantime, here are some useful articles:

* [Errors are values](https://go.dev/blog/errors-are-values)
* [Working with Errors](https://go.dev/blog/go1.13-errors) - Go 1.13 introduced some additional functionality in relation to errors



