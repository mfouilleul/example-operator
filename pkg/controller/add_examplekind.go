package controller

import (
	"github.com/mfouilleul/example-operator/pkg/controller/examplekind"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, examplekind.Add)
}