package command

import (
	"fmt"

	"github.com/multipleton/sa-4/engine"
)

type Print struct {
	Value string
}

func (p *Print) Execute(loop engine.Handler) {
	fmt.Println(p.Value)
}
