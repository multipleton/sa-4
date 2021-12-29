package parser

import (
	"testing"
	"fmt"
	
	"github.com/multipleton/sa-4/engine"
)

var cntRes engine.Command
func BenchmarkCount(b *testing.B) {
  const baseLen = 100
  line := "print line\n"
  for i := 0; i < 20; i++ {
    l := baseLen * (i + 1)
    for k := 0; k < l; k++ {
      line = line + "print line \n split a-b-c -\n"
    }
    b.Run(fmt.Sprintf("len=%d", l), func(b *testing.B) {
      cntRes = parse(line)
    })  
  }
}