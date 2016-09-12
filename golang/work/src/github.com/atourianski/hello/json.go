package main

import (
	"encoding/json"
	"fmt"
)

type box struct {
	width  int
	height int
	color  string
	open   bool
}

func main() {

	littleBox := box{width: 10, height: 15, color: "pink", open: true}

	b, _ := json.Marshal(littleBox)

	s := string(b)

	fmt.Println(s) //output is {} ...why?
}
