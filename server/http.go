package server

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Weather struct {
	Temperature string
}

func ServeMain(w http.ResponseWriter, r *http.Request) {
	temp := GetTemperature()
	model := Weather{Temperature: temp}
	dir, _ := os.Getwd()
	fmt.Printf("dir: " + dir)

	t, err := template.ParseFiles(dir+"/templates/header.tmpl", dir+"/templates/settings.tmpl", dir+"/templates/footer.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "settings", &model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
