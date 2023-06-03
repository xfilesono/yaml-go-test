package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fileName := "./config.yaml"
	var config a.Config
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(source))
}
