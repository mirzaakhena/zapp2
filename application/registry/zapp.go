package registry

import (
	"context"
	"flag"
	"fmt"
	"github.com/mirzaakhena/zapp2/application"
	"github.com/mirzaakhena/zapp2/controller/cmd"
	"github.com/mirzaakhena/zapp2/gateway/local"
	"github.com/mirzaakhena/zapp2/infrastructure/log"
	"github.com/mirzaakhena/zapp2/usecase/generateendtoend"
	"os"
)

type zapp struct {
  CommandMap    map[string]func(...string) error
  cmdController cmd.Controller
}

func (r *zapp) RunApplication() {

	fmt.Println("hello")

	flag.Parse()
	cmd := flag.Arg(0)

	if cmd == "" {
		log.Error(context.Background(), "zapp backend/frontend")
		return
	}

	var values = make([]string, 0)
	if flag.NArg() > 1 {
		values = flag.Args()[1:]
	}

	f, exists := r.CommandMap[cmd]
	if !exists {
		log.Error(context.Background(), "Command %s is not recognized", cmd)
		return
	}
	err := f(values...)
	if err != nil {
		log.Error(context.Background(), err.Error())
		return
	}

}

func NewZapp() func() application.RegistryContract {
  return func() application.RegistryContract {

    commandMap := make(map[string]func(...string) error, 0)

    datasource, err := local.NewLocalGateway()
    if err != nil {
      log.Error(context.Background(), "%v", err.Error())
      os.Exit(1)
    }

    return &zapp{
      CommandMap: commandMap,
      cmdController: cmd.Controller{
				CommandMap:             commandMap,
				GenerateEndToEndInport: generateendtoend.NewUsecase(datasource),
			},
    }

  }
}

func (r *zapp) SetupController() {
  r.cmdController.RegisterRouter()
}
