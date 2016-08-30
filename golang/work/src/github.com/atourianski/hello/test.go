package main 

import "fmt"

import "strconv"

func main(){

convertToC(50)

convertToM(50)

loops()

moreLoops()

fizzbuzz()

slices()

typesTest(5, "birch")

}

func convertToC(x int){

celsius := (x - 32) * 5/9

fmt.Println(string(celsius) + "C") 

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

func slices(){

x := []string { "grass", "fly", "plant" }

xx := append(x, "cow")

fmt.Println(xx)

slice1 := []int{1,2,3}

slice2 := make([]int, 2)

copy(slice2, slice1)

fmt.Println(slice1, slice2)

names := map[int]string{
1: "aaa",
2: "bbb",
3: "ccc",
4: "ddd",
5: "eee",
6: "fff",

}

www := [6]string{"a","b","c","d","e","f"}

slice3 := www[2:5]

fmt.Println("c d e")
fmt.Println(slice3)

fmt.Println(names[6], names[1])
fmt.Println(slice1[0])

}

func typesTest (x int, y string){

fmt.Println(strconv.Itoa(x) + " is a number and " + y + " is a string")
}




