package main

import (
	"github.com/boundedinfinity/bounded_xdg"
)

func main() {
	_, err := bounded_xdg.NewBoundeManager()

	if err != nil {
		panic(err)
	}
}
