package main 

import "fmt"


type person struct{

name string		
age int			//name, age and sex are all called fields 
sex string

}

type rectangle struct{

length, width int

}

func area(r1 *rectangle) int{

	return r1.length * r1.width
}

func main() {

	Alena := person{name:"Alena", age:21, sex:"Female"}	//created an instance of person

	Julia := person{name: "Julia", age: 26, sex: "Female"}	//created another instance of person

	fmt.Println(Alena, Julia.age)

	r1 := rectangle{length:5, width:10}

	fmt.Println("The area is ", area(r1))	//cannot use r1 (type rectangle) as type *rectangle						    //in argument to area???

}
