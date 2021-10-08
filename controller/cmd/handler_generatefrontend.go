package cmd

import (
  "context"
  "github.com/mirzaakhena/zapp2/infrastructure/log"
  "github.com/mirzaakhena/zapp2/infrastructure/util"
  "github.com/mirzaakhena/zapp2/usecase/generatefrontend"
)

// generateFrontendHandler ...
func (r *Controller) generateFrontendHandler(inputPort generatefrontend.Inport) func(...string) error {

  return func(...string) error {

    ctx := log.Context(context.Background())

    var req generatefrontend.InportRequest

    log.Info(ctx, util.MustJSON(req))

    _, err := inputPort.Execute(ctx, req)
    if err != nil {
      log.Error(ctx, err.Error())
    }

    return nil

  }
}
