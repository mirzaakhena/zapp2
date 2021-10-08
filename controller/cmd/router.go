package cmd

import (
	"github.com/mirzaakhena/zapp2/usecase/generateendtoend"
)

type Controller struct {
	CommandMap             map[string]func(...string) error
	GenerateEndToEndInport generateendtoend.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.CommandMap["frontend"] = r.generateEndToEndHandler(r.GenerateEndToEndInport)
}
