package main

import "fmt"

func main() {
	fmt.Println("Hello Danger!")

	for _, v := range []string{"foo", "bar", "baz"} {
		fmt.Println(v)
	}
}
