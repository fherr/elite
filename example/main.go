package main

import (
	"fmt"

	"github.com/fherr/elite"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", elite.GetRandomName())
	}
}
