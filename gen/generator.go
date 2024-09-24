package main

import (
	"bytes"
	"io"
	"path/filepath"
	"text/template"
)

type GenerateFile struct {
	FullFileName string
	Content      string
}

func GenerateByTemplate(t *template.Template, params *TemplateContextDTO) (string, error) {
	var b bytes.Buffer
	_ = io.Writer(&b)
	err := t.Execute(io.Writer(&b), params)
	return b.String(), err
}
func Generate(params *TemplateContextDTO, templateModel FileDeclaration) (*GenerateFile, error) {
	var b bytes.Buffer
	_ = io.Writer(&b)
	contentTemplate := template.Must(template.New(templateModel.Filename).Parse(templateModel.Template))
	content, err := GenerateByTemplate(contentTemplate, params)
	if err != nil {
		return nil, err
	}
	fileNameTemplate := template.Must(template.New(templateModel.Filename + "-filename").Parse(templateModel.Filename))
	fileName, err := GenerateByTemplate(fileNameTemplate, params)
	if err != nil {
		return nil, err
	}
	return &GenerateFile{
		FullFileName: filepath.Join(params.FilesDirectory, templateModel.RelativeDirectory, fileName),
		Content:      content,
	}, nil
}

func GenerateBulk(params *TemplateContextDTO, templates []FileDeclaration) ([]GenerateFile, error) {
	generatedFiles := make([]GenerateFile, 0)
	for _, templateModel := range templates {
		generateFile, err := Generate(params, templateModel)
		if err != nil {
			return generatedFiles, err
		}
		generatedFiles = append(generatedFiles, *generateFile)
	}
	return generatedFiles, nil
}
