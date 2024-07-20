package main

var (
	goMod = `module %s

go %s

require github.com/BrunProgramming/goxt v0.0.6-alpha

require (
	github.com/aymerick/raymond v2.0.2+incompatible // indirect
	github.com/bytedance/sonic v1.11.9 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.4 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.10.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)`

	mainGo = `package main

import (
    "github.com/BrunProgramming/goxt"
    "fmt"
)

func main() {
    router := goxt.NewRouter()
    router.Static("/style","./style")
    router.Get("/",func(c goxt.Ctx) {
        c.View("main",goxt.HbsCtx{},"")
    })
    router.Get("/:name",func(c goxt.Ctx) {
        name := c.Param("name")
        c.View("hello",goxt.HbsCtx{
            "name":name,
        },"")
    })
    fmt.Println("Server listening in http://localhost:8080")
    router.Run(":8080")
}`

	mainView = `<!DOCTYPE html>
<html>
  <head>
    <title>Goxt the best framework buffalo is noob</title>
    <link rel="stylesheet" href="/style/App.css">
  </head>
  <body>
   <h1>Plis enter your name below:</h1>
   <form id="form">
      <input type="text" id="input" autocomplete="off">
   </form>
    <script type="module">
      //midutrick copyright midudev©
      const $ = selector => document.querySelector(selector)
      
      $("#form").addEventListener("submit",e => {
        e.preventDefault()
        const value = $("#input").value
        const link = document.createElement("a")
        link.href=%s
        link.click()
      })
    </script>
  </body>
</html>`

	helloView = `<!DOCTYPE html>
<html>
  <head>
    <link rel="stylesheet" href="/style/App.css">
    <title>Goxt the best framework buffalo is noob</title>
  </head>
  <body>
   <h1>Hello {{name}}</h1>
  </body>
</html>`

	appCss = `body {
  text-align:center;
  background-color:#161B33;
  color:#fff;
  font-family:sans-serif;
}`
)
