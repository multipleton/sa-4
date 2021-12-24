package command

import (
	"fmt"

	"github.com/multipleton/sa-4/engine"
)

type Print struct {
	arg string
}

func (p *Print) Execute(loop *engine.Handler) {
	fmt.Println(p.arg)
}
