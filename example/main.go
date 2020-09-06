package main

import (
	"fmt"
	"io/ioutil"

	"github.com/binacsgo/pastebin"
)

func main() {
	bytes, err := ioutil.ReadFile("../README.md")
	if err != nil {
		panic(err)
	}
	output, _ := pastebin.ParseContentMD(string(bytes))
	fmt.Println(output)
}
