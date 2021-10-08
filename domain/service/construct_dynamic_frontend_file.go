package service

import (
  "fmt"
  "github.com/mirzaakhena/zapp2/domain/entity"
  "github.com/mirzaakhena/zapp2/infrastructure/templates"
  "os"
  "strings"
)

func createDynamicFrontendFile(tp *entity.ThePackage, et *entity.TheEntity) {

  // create store modules
  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/store/modules/file._js")
    outputFile := fmt.Sprintf("webapp/src/store/modules/%s.js", strings.ToLower(et.Name))
    basic(tp, templateFile, outputFile, et, 0664)
  }

  // create router modules
  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/router/modules/file._js")
    outputFile := fmt.Sprintf("webapp/src/router/modules/%s.js", strings.ToLower(et.Name))
    basic(tp, templateFile, outputFile, et, 0664)
  }
  {
    // create folder pages
    f := fmt.Sprintf("webapp/src/pages/%s", strings.ToLower(et.Name))
    os.Mkdir(f, 0777)

    // create file table under folder
    {
      templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/pages/folder/list._vue")
      outputFile := fmt.Sprintf("webapp/src/pages/%s/list.vue", strings.ToLower(et.Name))
      basic(tp, templateFile, outputFile, et, 0664)
    }

    // create file input under folder
    {
      templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/pages/folder/input._vue")
      outputFile := fmt.Sprintf("webapp/src/pages/%s/input.vue", strings.ToLower(et.Name))
      basic(tp, templateFile, outputFile, et, 0664)
    }

  }

}
