package service

import (
  "fmt"
  "github.com/mirzaakhena/zapp2/domain/entity"
  "os"
)

func makeFrontendDirectory(tp *entity.ThePackage) {

  var dir string

  dir = fmt.Sprintf("webapp/public")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("webapp/src/assets")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("webapp/src/store/modules")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("webapp/src/router/modules")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("webapp/src/pages")
  os.Mkdir(dir, 0777)

  dir = fmt.Sprintf("webapp/dist")
  os.Mkdir(dir, 0777)

  dir = fmt.Sprintf("webapp/src/utils")
  os.Mkdir(dir, 0777)

}
