package main 

import "fmt"

func main(){

convertToC(50)

convertToM(50)

loops()

moreLoops()

fizzbuzz()
}

func convertToC(x int){

celsius := (x - 32) * 5/9

fmt.Println(celsius)

}

func convertToM(f float64){

var metres float64 = f *  0.3048

fmt.Println(metres)

}

func loops(){

for i:= 1; i<10; i++ {

	if i % 2 == 0{
	fmt.Println("even")

	}else{
	fmt.Println("odd")}
}


}

func moreLoops(){

for i := 6; i >= 1; i--{

switch i {
	case 6: fmt.Println(i, "six")
	case 5: fmt.Println(i, "five")
	case 4: fmt.Println(i, "four")
	case 3: fmt.Println(i, "three")
	case 2: fmt.Println(i, "two")
	case 1: fmt.Println(i, "one")
	default: fmt.Println(i, "unknown number")
	}
}
}

func fizzbuzz(){

for i:=1; i<101; i++ {

	if i%3 == 0 && i%5 == 0 {
        fmt.Println("FizzBuzz")

	}else if i%3 == 0 {
	fmt.Println("Fizz")

	}else if i%5 == 0{
	fmt.Println("Buzz")

	}else{
	fmt.Println(i)
	}
}
}
