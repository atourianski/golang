package main

import "fmt"

func main() {
    fmt.Println("hello world")

age := 21

randNum := 5

fmt.Println(age - randNum)

myName := "Alena"

fmt.Println("My name is " + myName)

fmt.Printf("%x \nn", 100)

i := 1

for i < 11 {

	fmt.Println(i)

	i++
}

for j := 12; j > 5; j-- {

	if j == 12{
		fmt.Println("twelve")
	}else{
		fmt.Println(j)}
}

var favNums[5] float64

favNums[0] = 0.25
favNums[1] = 1.13
favNums[2] = 57.9
favNums[4] = 5.54

fmt.Println(favNums[2])

x := 6; x += 1

fmt.Println(x)

fmt.Println(32132*42452)

place := "first "
fmt.Println(place)

var (

 a = 2
 b = 10
 c = "string"
 d = false
)

another := "another variable"
var yet int
yet = 44
fmt.Println(another, yet)

fmt.Println(a, b, c, d)



}


	
