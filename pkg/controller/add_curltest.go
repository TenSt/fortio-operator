package controller

import (
	"github.com/TenSt/fortio-operator/pkg/controller/curltest"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, curltest.Add)
}
