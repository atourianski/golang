package main

import (
	"encoding/json"
	"fmt"
)

//json.Marshall takes a struct instance and converts it to json data (encodes it)

type Box struct {
	Width  int //make sure all the names are uppercase, or they won't be exported
	Height int
	Color  string
	Open   bool
}

func ExampleMarshal() {

	LittleBox := Box{Width: 10, Height: 15, Color: "pink", Open: true} //creating an instance of Box

	b, _ := json.Marshal(LittleBox) //encoding data within LittleBox and storing it within var b

	s := string(b) //converting the current type of uint8 to the type string

	fmt.Println(s)
}

//json.Unmarshal takes json data and converts it back to go code (decodes it)

type Result struct {
	Width  int
	Height int //must have a struct with matching fields for json.Unmarshal to work
	Color  string
	Open   bool
}

func main() {

	ExampleMarshal()

	text := []byte(`{"Width":10,"Height":15,"Color":"pink","Open":true}`) //<--raw json data

	var result Result

	json.Unmarshal(text, &result) //decoding the data within the var text and storing it in the struct Result found within the var result

	fmt.Println(result)

}
