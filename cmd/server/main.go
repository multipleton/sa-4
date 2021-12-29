package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"parser"
	"github.com/multipleton/sa-4/parser"
	//"github.com/multipleton/sa-4/command"
	"github.com/multipleton/sa-4/engine"
)

func main() {
	inputFile := os.Args[1]
	if _, err := os.Stat(inputFile); errors.Is(err, os.ErrNotExist) {
		fmt.Println("invalid file name.")
		return
	}
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parser.parse(commandLine)
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}

