// copyright 2016 CreatorQsF, de-liKeR
// MIT License
// `cliset` package provides some useful tools for developing cli tool in Golang.
package cliset

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/Qs-F/coloring"
	"golang.org/x/crypto/ssh/terminal"
)

type InMessage struct {
	Pipe string
	Args []string
}

type OutMessage struct {
	Body []string
}

func Warn(s string) {
	log.Println(coloring.Red(s))
}

func IsPiped() bool {
	if terminal.IsTerminal(0) { // pipe value is empty
		return false
	} else {
		return true
	}
}

func PipeValue() (s string, err error) {
	if IsPiped() {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		// if string(s) == "" {
		// 	return "", errors.New("Value is sent through pipe, but data is empty.")
		// }
		return string(b), nil
	} else {
		return "", errors.New("No value is sent through pipe.")
	}
}

func GetMsg() *InMessage {
	msg := &InMessage{}
	if IsPiped() {
		s, _ := PipeValue()
		msg.Pipe = s
	}
	if len(os.Args) > 1 {
		msg.Args = os.Args[1:]
	}
	return msg
}

// Subcmd returns 2 values, first type is bool and second one is int.
// First return value will be true when given subcommand name in argument is setted.
// Second value returns the number of nth element of `os.Args` which is value of builtin package `os`
// If there is no matched subcommand, second value will be 0.
// So first you need to check first bool value.
// func Subcmd(flagName string) (bool, int) {
// 	args := []string{}
// 	if len(os.Args) > 1 {
// 		args = os.Args
// 	} else {
// 		return false, 0
// 	}
// 	var wg sync.WaitGroup
// 	pos := make(chan int, len(args))
// 	wg.Add(len(args))
// 	for i, v := range args {
// 		go func(n int, s string) {
// 			defer wg.Done()
// 			if s == flagName {
// 				pos <- n
// 			}
// 		}(i, v)
// 	}
// 	wg.Wait()
// 	select {
// 	case n := <-pos:
// 		fmt.Printf("%v", <-pos)
// 		return true, n
// 	default:
// 		return false, 0
// 	}
// }
//

// RegSubcmd register your subcommand as alias of builtin flag package's flag.
// Before run `flag.Parse()`, RegSubcmd must be put.
// To use, 2 arguments needed.
// Set subcommand name in first argumanet.
// Set flag package's flag name last argument.
func RegSubcmd(subcmd, flagName string) {
	if !(len(os.Args) > 1) {
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(os.Args))
	for i, v := range os.Args {
		go func(n int, s string) {
			if s == subcmd {
				os.Args[n] = "-" + flagName
			}
			wg.Done()
		}(i, v)
	}
	wg.Wait()
}

// Low

func LowSubcmd(subcmd string) (bool, int) {
	args := []string{}
	if len(os.Args) > 1 {
		args = os.Args
	} else {
		return false, 0
	}
	for i, v := range args {
		if v == subcmd {
			return true, i
		}
	}
	return false, 0
}

func LowRegSubcmd(subcmd, flagName string) {
	b, n := LowSubcmd(subcmd)
	if b {
		os.Args[n] = "-" + flagName
	}
}
