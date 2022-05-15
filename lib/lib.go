package lib

import (
	"github.com/robertkrimen/otto"
)

func RunCode(input string) (*otto.Value, error) {
	vm := otto.New()
	value, err := vm.Eval(input)
	if err != nil {
		return nil, err
	}
	return &value, err
}
