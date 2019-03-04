package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func Run(args []string) error {
	if len(args) != 2 {
		return errors.New("pow X Y\nPrints X to the power of Y")
	}
	x, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return err
	}
	y, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return err
	}
	fmt.Println(math.Pow(x, y))
	return nil
}

func main() {}
