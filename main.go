package main

import (
  "runtime"
  "os"
  "fmt"
  "strings"
)

type ArgFunc func(arg string)
type Args map[string]ArgFunc

func ArgvCompile(args Args) {
  for k, v := range os.Args {
    if args[v] != nil {
      args[v](os.Args[k+1])
    }
  }
}

func CreateFiles(name string)  {
  main,err := os.Create(fmt.Sprintf("%s/main.go",name))
  if err != nil {
    panic(err)
  }
  main.WriteString(mainGo)
  os.Mkdir(fmt.Sprintf("%s/views",name),0700)
  os.Mkdir(fmt.Sprintf("%s/style",name),0700)
  mainHbs,_ := os.Create(fmt.Sprintf("%s/views/main.hbs",name))

  mainHbs.WriteString(fmt.Sprintf(mainView,"`${location.href}${value}`"))
  helloHbs,_ := os.Create(fmt.Sprintf("%s/views/hello.hbs",name))
  
  helloHbs.WriteString(helloView)
  
  style,_ := os.Create(fmt.Sprintf("%s/style/App.css",name))
  style.WriteString(appCss)
}

func main() {
  var name,repo string

  args := Args{
    "-name":func(arg string) {
      name=arg
      os.Mkdir(name,0700)
    },
    "-repo":func(arg string) {
      repo=arg
      modFile,_ := os.Create(fmt.Sprintf("%s/go.mod",name))
      modFile.WriteString(fmt.Sprintf(goMod,repo,strings.Replace(runtime.Version(),"go","",1)))
    },
  }
  ArgvCompile(args)
  CreateFiles(name)
  fmt.Printf(`Now start hacking
      cd %s
      and execute go get for install the modules
      for exec the app use go run .
  `,name)
}
