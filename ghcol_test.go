package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetCodeInput(t *testing.T) {
	a := assert.New(t)
	simpleInput := []byte("def helloWorld\n\tputs 'hello, world'\nend\n^")
	complexInput := []byte(`nc·006() /usr/local/Cellar/go/1.3.1/libexec/src/pkg/
    testing/testing.go:416
    +0x176 runtime.panic(0x2d9440, 0x4b8e2f) /usr/local/Cellar/go/1.3.1/libexec/
    src/pkg/runtime/panic.c:248 +0x18d _/Volumes/Code/dring1-github/ghco`)

	r := bytes.NewBuffer(simpleInput)
	code, _ := ReadInput(r)
	a.Equal(FormatString(code), "def+helloWorld+puts+%27hello%2C+world%27%0Aend%0A%5E")

	r = bytes.NewBuffer(complexInput)
	code, _ = ReadInput(r)
	// a.Equal(FormatString(code), `nc%C2%B7006%28%29+%2Fusr%2Flocal%2FCellar%2Fgo%2F
	//    1.3.1%2Flibexec%2Fsrc%2Fpkg%2F%0A++++testing%2Ftesting.go%3A416%0A++++%2B0x176+runtime.panic%280x2d9440%2C+0x4b`)
}
