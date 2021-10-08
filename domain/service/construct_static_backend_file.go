package service

import (
  "fmt"
  "github.com/mirzaakhena/zapp2/domain/entity"
  "github.com/mirzaakhena/zapp2/infrastructure/templates"
)

func createStaticBackendFile(tp *entity.ThePackage) {
  // backend
  {
    {
      templateFile, _ := templates.Templates.ReadFile("backend/main._go")
      outputFile := fmt.Sprintf("main.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/model/model_basic._go")
      outputFile := fmt.Sprintf("model/system-model.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/model/user-model._go")
      outputFile := fmt.Sprintf("model/system-user.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/repository/user-repository._go")
      outputFile := fmt.Sprintf("repository/system-user.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/service/auth/user-service._go")
      outputFile := fmt.Sprintf("service/auth/system-user.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/controller/restapi/user-controller._go")
      outputFile := fmt.Sprintf("controller/restapi/system-user.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/controller/restapi/router._go")
      outputFile := fmt.Sprintf("controller/restapi/system-router.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/utils/identifier._go")
      outputFile := fmt.Sprintf("shared/utils/identifier.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/utils/password._go")
      outputFile := fmt.Sprintf("shared/utils/password.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/utils/json._go")
      outputFile := fmt.Sprintf("shared/utils/json.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/utils/strings._go")
      outputFile := fmt.Sprintf("shared/utils/strings.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/token/contract._go")
      outputFile := fmt.Sprintf("shared/token/contract.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/token/implementation._go")
      outputFile := fmt.Sprintf("shared/token/implementation.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/token/public._go")
      outputFile := fmt.Sprintf("shared/token/public.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/config/contract._go")
      outputFile := fmt.Sprintf("shared/config/contract.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/config/implementation._go")
      outputFile := fmt.Sprintf("shared/config/implementation.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/config/public._go")
      outputFile := fmt.Sprintf("shared/config/public.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/log/contract._go")
      outputFile := fmt.Sprintf("shared/log/contract.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/log/implementation._go")
      outputFile := fmt.Sprintf("shared/log/implementation.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/log/public._go")
      outputFile := fmt.Sprintf("shared/log/public.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/transaction/contract._go")
      outputFile := fmt.Sprintf("shared/transaction/contract.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/transaction/implementation._go")
      outputFile := fmt.Sprintf("shared/transaction/implementation.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/transaction/public._go")
      outputFile := fmt.Sprintf("shared/transaction/public.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/messagebroker/implementation-producer._go")
      outputFile := fmt.Sprintf("shared/messagebroker/implementation-producer.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/messagebroker/implementation-consumer._go")
      outputFile := fmt.Sprintf("shared/messagebroker/implementation-consumer.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/messagebroker/contract._go")
      outputFile := fmt.Sprintf("shared/messagebroker/contract.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/messagebroker/public._go")
      outputFile := fmt.Sprintf("shared/messagebroker/public.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    //{
    //  templateFile, _ := templates.Templates.ReadFile("backend/shared/httpclient/contract._go")
    //  outputFile := fmt.Sprintf("shared/httpclient/contract.go")
    //  basic(tp, templateFile, outputFile, tp, 0664)
    //}
    //
    //{
    //  templateFile, _ := templates.Templates.ReadFile("backend/shared/httpclient/implementation._go")
    //  outputFile := fmt.Sprintf("shared/httpclient/implementation.go")
    //  basic(tp, templateFile, outputFile, tp, 0664)
    //}
    //
    //{
    //  templateFile, _ := templates.Templates.ReadFile("backend/shared/httpclient/public._go")
    //  outputFile := fmt.Sprintf("shared/httpclient/public.go")
    //  basic(tp, templateFile, outputFile, tp, 0664)
    //}

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/error/error_code._go")
      outputFile := fmt.Sprintf("shared/error/error_code.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/error/error._go")
      outputFile := fmt.Sprintf("shared/error/error.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/error/error_test._go")
      outputFile := fmt.Sprintf("shared/error/error_test.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/constant/constant._go")
      outputFile := fmt.Sprintf("shared/constant/constant.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/shared/extractor/extractor._go")
      outputFile := fmt.Sprintf("shared/extractor/extractor.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/._gitignore")
      outputFile := fmt.Sprintf(".gitignore")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/config._toml")
      outputFile := fmt.Sprintf("config.toml")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/build._sh")
      outputFile := fmt.Sprintf("build.sh")
      basic(tp, templateFile, outputFile, tp, 0755)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/README._md")
      outputFile := fmt.Sprintf("README.md")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/go._mod")
      outputFile := fmt.Sprintf("go.mod")
      basic(tp, templateFile, outputFile, tp, 0664)
    }

    {
      templateFile, _ := templates.Templates.ReadFile("backend/webapp/webapp._go")
      outputFile := fmt.Sprintf("webapp/webapp.go")
      basic(tp, templateFile, outputFile, tp, 0664)
    }
  }
}
