package play003_error

import (
	"errors"
	"fmt"
)

type TypicalErr struct {
	e string
}

func (t TypicalErr) Error() string {
	return t.e
}

func Test() {
	err := TypicalErr{"typical error"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	var e TypicalErr
	if errors.As(err2, &e) {
		println("TypicalErr is on the chain of err2")
		println(err == e)
	} else {
		println("TypicalErr is not on the chain of err2")
	}
}
