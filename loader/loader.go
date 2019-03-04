package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("loader says hello!")
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
