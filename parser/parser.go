package parser

import (
	"strings"
	
	"github.com/multipleton/sa-4/command"
	"github.com/multipleton/sa-4/engine"
)

func parse(line string) engine.Command {
	arr := strings.Fields(line)
	if arr[0] == "print" {
		if len(arr) == 2 {
			return &command.Print{Value: arr[1]}
		}
		return &command.Print{Value: "SYNTAX ERROR: invalid arguments count"}
	}
	if arr[0] == "split" {
		if len(arr) == 3 {
			if len(arr[2]) == 1 {
				r := []rune(arr[2])
				return &command.Split{Input: arr[1], Separator: r[0]}
			}
			return &command.Print{Value: "SYNTAX ERROR: separator should be a single character"}
		}
		return &command.Print{Value: "SYNTAX ERROR: invalid arguments count"}
	}
	return &command.Print{Value: "SYNTAX ERROR: invalid command"}
}
