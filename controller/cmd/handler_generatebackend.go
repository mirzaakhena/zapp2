package cmd

import (
  "context"
  "github.com/mirzaakhena/zapp2/infrastructure/log"
  "github.com/mirzaakhena/zapp2/infrastructure/util"
  "github.com/mirzaakhena/zapp2/usecase/generatebackend"
)

// generateBackendHandler ...
func (r *Controller) generateBackendHandler(inputPort generatebackend.Inport) func(...string) error {

  return func(...string) error {

    ctx := log.Context(context.Background())

    var req generatebackend.InportRequest

    log.Info(ctx, util.MustJSON(req))

    _, err := inputPort.Execute(ctx, req)
    if err != nil {
      log.Error(ctx, err.Error())
    }

    return nil

  }
}
