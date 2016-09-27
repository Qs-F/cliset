package main

import (
	"flag"
	"fmt"
	"github.com/Qs-F/cliset"
)

func main() {
	g := flag.Int("n", 0, "hoge")
	flag.Parse()
	b, _ := cliset.Subcmd("init")
	fmt.Println(*g)
	if b {
		fmt.Println("init?")
	}
}
