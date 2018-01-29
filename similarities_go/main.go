package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"templating"
)

type Configuration struct {
	LayoutPath  string
	IncludePath string
}

func loadConfiguration(fileName string) {
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("error:", err)
	}
	log.Println("layout path: ", configuration.LayoutPath)
	log.Println("include path: ", configuration.IncludePath)
	templating.SetTemplateConfig(configuration.LayoutPath, configuration.IncludePath)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var g [2]string
	if r.Method == http.MethodPost {
		firstString := r.FormValue("firstString")
		secondString := r.FormValue("secondString")
		g = [2]string{firstString, secondString}
	}

	templating.RenderTemplate(w, "index.tmpl", g)
}

func main() {
	loadConfiguration("config.json")
	templating.LoadTemplates()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
