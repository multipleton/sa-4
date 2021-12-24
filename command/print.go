package command

import (
	"fmt"

	"github.com/multipleton/sa-4/engine"
)

type Print struct {
	value string
}

func (p *Print) Execute(loop engine.Handler) {
	fmt.Println(p.value)
}
