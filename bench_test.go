package parser

import (
	"testing"
	"fmt"
	
	"github.com/multipleton/sa-4/engine"
	"github.com/multipleton/sa-4/parser"
)

var cntRes engine.Command
func BenchmarkCount(b *testing.B) {
  const baseLen = 2000
  line := "print line\n"
  for i := 0; i < 14; i++ {
    l := baseLen * (i + 1)
    for k := 0; k < l; k++ {
      line = line + "print line \nsplit a-b-c -\n"
    }
    b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
      cntRes = parser.Parse(line)
    })  
  }
}
