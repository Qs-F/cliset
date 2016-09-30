package main

import (
	"flag"
	"fmt"
	"github.com/Qs-F/cliset"
)

func main() {
	// cliset.LowRegSubcmd("point", "p")
	// cliset.LowRegSubcmd("test", "p")
	// p := flag.Bool("p", false, "hoge")
	// g := flag.Int("n", 0, "hoge")
	// flag.Parse()
	// b, _ := cliset.LowSubcmd("init")
	// fmt.Println(*g)
	// fmt.Println(*p)
	// if b {
	// 	fmt.Println("init?")
	// }
	//
	cliset.RegSubcmd("fizz", "f")
	cliset.RegSubcmd("Buzz", "b")
	cliset.RegSubcmd("b", "b")
	f := flag.Int("f", 0, "Fizz number")
	b := flag.Bool("b", false, "Buzz true or false")
	flag.Parse()
	fmt.Println("f is ", *f)
	fmt.Println("b is ", *b)
}
