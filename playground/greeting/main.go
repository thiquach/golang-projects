// Functions as first-class citizens: assigned to variables,
// as arguments to other functions, returned from other functions.
package main

import "fmt"

func greeting(name string, formatter func(string) string) {
	fmt.Println(formatter(name))
}

func formalGreeting(name string) string {
	return "Hello Mr/Ms " + name
}

func main() {
	casualGreeting := func(name string) string {
		return "Hey " + name + "!"
	}

	greeting("Smith", formalGreeting)
	greeting("Sam", casualGreeting)
}