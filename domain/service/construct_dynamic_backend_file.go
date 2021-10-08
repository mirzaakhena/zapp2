package service

import (
  "fmt"
  "github.com/mirzaakhena/zapp2/domain/entity"
  "github.com/mirzaakhena/zapp2/infrastructure/templates"
  "strings"
)

func createDynamicBackendFile(tp *entity.ThePackage, et *entity.TheEntity) {

  {
    templateFile, _ := templates.Templates.ReadFile("backend/repository/repository._go")
    outputFile := fmt.Sprintf("repository/%s.go", strings.ToLower(et.Name))
    basic(tp, templateFile, outputFile, et, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/controller/restapi/controller._go")
    outputFile := fmt.Sprintf("controller/restapi/%s.go", strings.ToLower(et.Name))
    basic(tp, templateFile, outputFile, et, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/model/model._go")
    outputFile := fmt.Sprintf("model/%s.go", strings.ToLower(et.Name))
    basic(tp, templateFile, outputFile, et, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/service/crud/service._go")
    outputFile := fmt.Sprintf("service/crud/%s.go", strings.ToLower(et.Name))
    basic(tp, templateFile, outputFile, et, 0664)
  }

}

func createEnumFile(tp *entity.ThePackage, en *entity.TheEnum) {
  {
    templateFile, _ := templates.Templates.ReadFile("backend/model/enum/enum._go")
    outputFile := fmt.Sprintf("model/enum/%s.go", strings.ToLower(en.Name))
    basic(tp, templateFile, outputFile, en, 0664)
  }
}
