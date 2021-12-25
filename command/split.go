package command

import (
	"github.com/multipleton/sa-4/engine"
)

type Split struct {
	Input     string
	Separator rune
}

func (s *Split) Execute(loop engine.Handler) {
	var splitted []string
	var buffer string
	for _, entry := range s.Input {
		if entry == s.Separator && len(buffer) > 0 {
			splitted = append(splitted, buffer)
			buffer = ""
			continue
		}
		buffer += string(entry)
	}
	splitted = append(splitted, buffer)
	for _, entry := range splitted {
		loop.Post(&Print{Value: entry})
	}
}
