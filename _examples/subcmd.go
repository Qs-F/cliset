package main

import (
	"flag"
	"fmt"
	"github.com/Qs-F/cliset"
)

func main() {
	cliset.RegSubcmd("fizz", "f")
	cliset.RegSubcmd("Buzz", "b")
	cliset.RegSubcmd("b", "b")
	f := flag.Int("f", 0, "Fizz number")
	b := flag.Bool("b", false, "Buzz true or false")
	flag.Parse()
	fmt.Println("f is ", *f)
	fmt.Println("b is ", *b)
}
