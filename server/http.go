package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

type Weather struct {
	Temperature string
}

var templates map[string]*template.Template

func InitTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	fmt.Printf("Creating")
	templatesDir := "/opt/weatherServer/"
	templs, err := filepath.Glob(templatesDir + "templates/*.tmpl")
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}

	for _, t := range templs {
		fmt.Printf("Parsing %s", t)
		templates[filepath.Base(t)] = template.Must(template.ParseFiles(templs...))
	}
}

func servePage(w http.ResponseWriter, r *http.Request, name string, data map[string]interface{}) {

	tmpl, ok := templates[name]
	if !ok {
		fmt.Printf("Template %s does not exist", name)
		return
	}
	fmt.Printf("Serving %s", name)
	tmpl.ExecuteTemplate(w, "base", data)
}
func ServePage(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	fmt.Printf("Try to serve")
	data = make(map[string]interface{})
	q := buildQuery(NOW, time.Now())
	temp := GetTemperature(q)
	data["Temperature"] = temp
	q = buildQuery(MIN, time.Now())
	temp = GetTemperature(q)
	data["MinTemp"] = temp
	q = buildQuery(MAX, time.Now())
	temp = GetTemperature(q)
	data["MaxTemp"] = temp

	servePage(w, r, "index.tmpl", data)
}
