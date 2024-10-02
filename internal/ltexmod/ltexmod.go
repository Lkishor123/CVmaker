package ltexmod

import (
	models "CVmaker/internal/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

func ParseJSON(filename string) (*models.Resume, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var resume models.Resume
	err = json.Unmarshal(data, &resume)
	if err != nil {
		return nil, err
	}
	return &resume, nil
}
func CompilePDF(texFile string) error {
	cmd := exec.Command("pdflatex", texFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func GenerateLaTeX(resume *models.Resume, templateFile, outputFile string) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, resume)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputFile, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
