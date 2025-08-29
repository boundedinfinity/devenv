package main

import (
	"github.com/boundedinfinity/bounded_xdg"
)

func main() {
	bm, err := bounded_xdg.NewBoundeManager()

	if err != nil {
		panic(err)
	}

	bounded_xdg.PP(bm.Shells())

	// if err := bm.Enabled("bash", "gnupg", true); err != nil {
	// 	panic(err)
	// }
}
