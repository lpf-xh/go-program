package main

import (
	"fmt"
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("./calc.so")
	if err != nil {
		log.Fatal(err)
	}

	add, err := p.Lookup("Add")
	if err != nil {
		log.Fatal(err)
	}

	sub, err := p.Lookup("Sub")
	if err != nil {
		log.Fatal(err)
	}

	sumRet := add.(func(int, int) int)(3, 2)
	fmt.Println(sumRet)

	subRet := sub.(func(int, int) int)(3, 2)
	fmt.Println(subRet)
}

// go build -buildmode=plugin -o calc.so calc.go
// go run main.go
