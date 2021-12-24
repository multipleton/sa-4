package command

import (
	"strings"

	"github.com/multipleton/sa-4/engine"
)

type Split struct {
	input     string
	separator string
}

func (s *Split) Execute(loop engine.Handler) {
	splitted := strings.Split(s.input, s.separator)
	for _, entry := range splitted {
		loop.Post(&Print{value: entry})
	}
}
