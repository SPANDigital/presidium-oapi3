package main

//go:generate go-bindata --nometadata -pkg tpl -o pkg/service/tpl/tpl.go pkg/templates/...

import (
	"github.com/SPANDigital/presidium-oapi3/cmd"
)

func main() {
	cmd.Execute()
}
