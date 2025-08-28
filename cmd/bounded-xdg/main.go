package main

import (
	"fmt"

	"github.com/boundedinfinity/bounded_xdg"
)

func main() {
	bm, err := bounded_xdg.NewBoundeManager()

	if err != nil {
		panic(err)
	}

	fmt.Println(bm.Shells())
}
