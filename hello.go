package main

import (
	"fmt"
	"runtime/debug"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())

	bi, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Println("path:", bi.Path, "main.path:", bi.Main.Path, "main.version:", bi.Main.Version, "main.sum:", bi.Main.Sum, "replaced:", bi.Main.Replace)
		for _, d := range bi.Deps {
			fmt.Println(d.Path, d.Version, d.Sum)
		}
	}

}
