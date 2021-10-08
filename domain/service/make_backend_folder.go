package service

import (
  "fmt"
  "github.com/mirzaakhena/zapp2/domain/entity"
  "os"
)

func makeBackendDirectory(tp *entity.ThePackage) {

  var dir string

  dir = fmt.Sprintf("controller/restapi")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("model")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("model/enum")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("service/auth")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("service/crud")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("repository")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/utils")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/token")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/config")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/log")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/transaction")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/messagebroker")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/httpclient")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/constant")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/extractor")
  os.MkdirAll(dir, 0777)

  dir = fmt.Sprintf("shared/error")
  os.MkdirAll(dir, 0777)

}
