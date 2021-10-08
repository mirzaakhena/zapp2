package service

import (
  "fmt"
  "github.com/mirzaakhena/zapp2/domain/entity"
  "github.com/mirzaakhena/zapp2/infrastructure/templates"
)

func createStaticFrontendFile(tp *entity.ThePackage) {
  // frontend
  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/public/favicon._ico")
    outputFile := fmt.Sprintf("webapp/public/favicon.ico")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/public/index._html")
    outputFile := fmt.Sprintf("webapp/public/index.html")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/babel.config._js")
    outputFile := fmt.Sprintf("webapp/babel.config.js")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/vue.config._js")
    outputFile := fmt.Sprintf("webapp/vue.config.js")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/package._json")
    outputFile := fmt.Sprintf("webapp/package.json")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/App._vue")
    outputFile := fmt.Sprintf("webapp/src/App.vue")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/assets/style._css")
    outputFile := fmt.Sprintf("webapp/src/assets/style.css")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/dist/index._html")
    outputFile := fmt.Sprintf("webapp/dist/index.html")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/pages/forgotpassword._vue")
    outputFile := fmt.Sprintf("webapp/src/pages/forgotpassword.vue")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/pages/home._vue")
    outputFile := fmt.Sprintf("webapp/src/pages/home.vue")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/pages/login._vue")
    outputFile := fmt.Sprintf("webapp/src/pages/login.vue")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/pages/notfound._vue")
    outputFile := fmt.Sprintf("webapp/src/pages/notfound.vue")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/pages/register._vue")
    outputFile := fmt.Sprintf("webapp/src/pages/register.vue")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/pages/successregister._vue")
    outputFile := fmt.Sprintf("webapp/src/pages/successregister.vue")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/main._js")
    outputFile := fmt.Sprintf("webapp/src/main.js")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/store/index._js")
    outputFile := fmt.Sprintf("webapp/src/store/index.js")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/store/crudtable._js")
    outputFile := fmt.Sprintf("webapp/src/store/crudtable.js")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/router/index._js")
    outputFile := fmt.Sprintf("webapp/src/router/index.js")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/utils/httprequest._js")
    outputFile := fmt.Sprintf("webapp/src/utils/httprequest.js")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/utils/auth._js")
    outputFile := fmt.Sprintf("webapp/src/utils/auth.js")
    basic(tp, templateFile, outputFile, tp, 0664)
  }

  {
    templateFile, _ := templates.Templates.ReadFile("backend/webapp/src/utils/filter._js")
    outputFile := fmt.Sprintf("webapp/src/utils/filter.js")
    basic(tp, templateFile, outputFile, tp, 0664)
  }
}
