package service

import (
	"bytes"
	"fmt"
	"github.com/SPANDigital/presidium-oapi3/pkg/service/tpl"
	"github.com/getkin/kin-openapi/openapi3"
	"os"
	"text/template"
)

type InfoService interface {
	ProcessInfo(info *openapi3.Info) error
}

type infoService struct {
	t *template.Template
}

func NewInfoService() (InfoService, error) {
	t, err := template.New("info.gomd").Funcs(tpl.FuncMap()).ParseFiles(
		"pkg/templates/info.gomd",
	)
	if err != nil {
		return &infoService{}, err
	}
	return &infoService{t}, nil
}

func (c *infoService) ProcessInfo(info *openapi3.Info) error {
	dir := "out/content/_reference/"
	path := fmt.Sprintf("%s/01_%s.md", dir, "info")
	os.MkdirAll(dir, os.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	err = c.t.Execute(&b, info)
	if err != nil {
		return err
	}
	cleaned := cleanForMarkdown(b)
	f.Write(cleaned.Bytes())
	return nil
}
