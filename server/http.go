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

var infServer InfluxServ

func SetServer(s InfluxServ) {
	infServer = s
}

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
	getTemp(&data, "Now", NOW)
	getTemp(&data, "Min", MIN)
	getTemp(&data, "Max", MAX)

	servePage(w, r, "index.tmpl", data)
}

func getTemp(d *map[string]interface{}, k string, t Temperature) {
	q := buildQuery(t, time.Now())
	temp, err := infServer.GetTemperature(q)
	if err != nil {

	} else {
		index := k + "Temp"
		(*d)[index] = temp.temp
		index = k + "Time"
		(*d)[index] = temp.time
	}
}
