package main

import(
	"fmt"
)

func main(){

add := closure()

fmt.Println(add())

//deferExample()

countDown(3)

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

//recursion is when a function loops on itself by calling itself within its own body, example:

func countDown(x int) int {

	if (x > 0) {
		 countDown(x-1)	//<-- countDown calls itself here
	}
	fmt.Println(x)
	return x	
}
	

//defer postpones the execution of a function until the surrounding function returns, example:

//func deferExample() string {

//	defer fmt.Println("World")

//	fmt.Println("Hello")
	
//}

//panic aborts if a function returns an error so that you don't have to deal with the broken function running all //the way through, example:



