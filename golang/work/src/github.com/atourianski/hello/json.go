package main

import (
	"encoding/json"
	"fmt"
)

type Box struct {
	Width  int //make sure all the names are uppercase, or they won't be exported
	Height int
	Color  string
	Open   bool
}

func ExampleMarshal() {
	LittleBox := Box{Width: 10, Height: 15, Color: "pink", Open: true}

	b, _ := json.Marshal(LittleBox) //json.Marshall takes a go instance and converts it to json data (encodes it)

	s := string(b)

	fmt.Println(s)

}

func main() {

	ExampleMarshal()

	//json.Unmarshal takes json data and converts it back to go code (decodes it)

	text := "[{\"Width\":10,\"Height\":15,\"Color\":\"pink\",\"Open\":true}]"

	bytes := []byte(text)

	x := []string{}

	json.Unmarshal(bytes, &x)

	fmt.Println(x)

}
