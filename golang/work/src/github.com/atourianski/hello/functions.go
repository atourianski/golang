package main

import(
	"fmt"
)

func main(){

add := closure()

fmt.Println(add())

deferExample()

}

// a closure is a function with preserved data, it has access to the parent scope, even after the
// parent function has closed, example:

func closure() func() int{
	
	x := 5
	inner := func() int{
	return x + 5
	}
	return inner 
}

//defer postpones the execution of a function until the surrounding function returns, example:

func deferExample() string {

	defer fmt.Println("World")

	fmt.Println("Hello")
	
}
