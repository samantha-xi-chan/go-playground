package play002_error

import (
	"errors"
	"fmt"
)

var BaseErr = errors.New("base error")

// BaseErr -> err1 -> err2
// err3
func Test() {
	err1 := fmt.Errorf("wrap base: %w", BaseErr)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	err3 := errors.New("base error")
	println(err2 == BaseErr)
	println(err2 == err3)

	if errors.Is(err1, BaseErr) {
		println("err1 is BaseErr")
	} else {
		println("err1 is not BaseErr")
	}

	if errors.Is(err2, BaseErr) {
		println("err2 is BaseErr")
	} else {
		println("err2 is not BaseErr")
	}

	if errors.Is(err3, BaseErr) {
		println("err3 is BaseErr")
	} else {
		println("err3 is not BaseErr")
	}
}
