package main

import "fmt"

func greeting(name string) string {
	return "Yo my man " + name
}

func printAge() int {
	return 22
}

func main() {
	fmt.Println(greeting("sukhmani"))
	fmt.Printf(greeting("sukhmani ")+"your age is %d", printAge())
}
