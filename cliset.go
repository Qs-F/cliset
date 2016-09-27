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

func Subcmd(flagName string) (bool, int) {
	args := []string{}
	if len(os.Args) > 1 {
		args = os.Args
	} else {
		return false, 0
	}
	var wg sync.WaitGroup
	pos := make(chan int, 1)
	wg.Add(len(args))
	for i, v := range args {
		go func(n int, s string) {
			defer wg.Done()
			if s == flagName {
				pos <- n
			}
		}(i, v)
	}
	wg.Wait()
	select {
	case n := <-pos:
		return true, n
	default:
		return false, 0
	}
}
