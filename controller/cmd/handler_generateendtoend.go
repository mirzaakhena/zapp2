package cmd

import (
  "context"
  "fmt"
  "github.com/mirzaakhena/zapp2/infrastructure/log"
  "github.com/mirzaakhena/zapp2/usecase/generateendtoend"
)

// generateEndToEndHandler ...
func (r *Controller) generateEndToEndHandler(inputPort generateendtoend.Inport) func(...string) error {

  return func(commands ...string) error {

    ctx := log.Context(context.Background())

    if len(commands) < 1 {
      err := fmt.Errorf("invalid input")
      return err
    }

    //var req gencontroller.InportRequest
    //req.ControllerName = commands[0]
    //req.UsecaseName = commands[1]

    var req generateendtoend.InportRequest
    req.YamlFile = commands[0]

    _, err := inputPort.Execute(ctx, req)
    if err != nil {
      log.Error(ctx, err.Error())
    }

    return nil

  }
}
