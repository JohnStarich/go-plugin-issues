package main

import "C"
import (
	"fmt"
	"plugin"
)

//export doit
func doit() {
	fmt.Println("doit() says hello!")
	p, err := plugin.Open("out/pow")
	if err != nil {
		panic(err)
	}

	sym, err := p.Lookup("Run")
	if err != nil {
		panic(err)
	}

	runFunc, ok := sym.(func([]string) error)
	if !ok {
		panic("invalid symbol for plugin Run func")
	}

	err = runFunc(nil)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}

func main() {}
