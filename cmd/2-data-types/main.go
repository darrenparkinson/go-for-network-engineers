package main

import "fmt"

func main() {
	// ints
	var i1 int
	i1 = 123
	i2 := 456
	fmt.Println(i1, i2)
	fmt.Printf("%T, %T\n", i1, i2)

	// floats
	var f1 float32
	f1 = -1.12
	f2 := 3.14159
	fmt.Println(f1, f2)
	fmt.Printf("%T, %T\n", f1, f2)

	// bool
	var b1 bool
	b1 = true
	b2 := false
	fmt.Println(b1, b2)
	fmt.Printf("%T, %T\n", b1, b2)

	// string
	var s1 string
	s1 = "hello"
	s2 := "world"
	fmt.Println(s1, s2)
	fmt.Printf("%T, %T\n", s1, s2)

	// bytes
	var t1 byte
	t1 = 0xff
	t2 := []byte("hello")
	fmt.Println(t1, t2)
	fmt.Printf("%T, %T\n", t1, t2)

}
