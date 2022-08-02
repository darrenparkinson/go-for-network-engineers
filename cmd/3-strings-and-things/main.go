package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// string initialisation
	var s1 string
	s1 = "hello"
	s2 := "world"
	fmt.Println(s1, s2) // note that Println inserts a space between each passed variable
	fmt.Printf("%T, %T\n", s1, s2)

	// inefficient string concatenation
	s3 := "Hello, " + "world!"
	fmt.Println(s3)

	// string builder
	var hello strings.Builder // remember to import strings
	hello.WriteString("Hello, ")
	hello.WriteString("world!")
	fmt.Println(hello.String())

	// formatted strings
	hundred := 100
	pi := 3.14
	result := fmt.Sprintf("One hundred: %d and Pi: %.2f", hundred, pi)
	fmt.Println(result)

	fmt.Printf("One hundred: %d and Pi: %.2f\n", hundred, pi)

	// useful fmt function
	s := "Bob is 195cm tall and weighs 89kg"
	var name string
	var height, weight int
	fmt.Sscanf(s, "%s is %dcm tall and weighs %dkg", &name, &height, &weight)
	fmt.Println(name, height, weight)

	// strings package
	fmt.Println(strings.HasPrefix("https://www.cisco.com", "http"))                // true
	fmt.Println(strings.HasSuffix("https://www.cisco.com", "COM"))                 // false, case sensitive
	fmt.Println(strings.ToLower("COM"))                                            // "com"
	fmt.Println(strings.Split("www.cisco.com", ","))                               // ["www","cisco","com"]
	fmt.Println(strings.Fields("   ip dhcp excluded-address 10.0.1.1 10.0.1.10 ")) // ["ip", "dhcp", "excluded-address", "10.0.1.1", "10.0.1.10"]
	fmt.Println(strings.Join([]string{"enable", "password", "cisco123"}, " "))
	fmt.Println("ba" + strings.Repeat("na", 2))

	// strconv package
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println(b, f, i, u)

	// errors
	b1, err := strconv.ParseBool("true")
	fmt.Println(b1, err)
	b2, err := strconv.ParseBool("networking")
	fmt.Println(b2, err)
	b3, _ := strconv.ParseBool("true")
	fmt.Println(b3)

}
