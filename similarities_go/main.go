package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("assets/templates/*.tmpl"))
}

func GetStrings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := tpl.ParseFiles("assets/templates/output.tmpl")
	if err != nil {
		log.Fatalln(err)
	}
	firstString := r.FormValue("firstString")
	secondString := r.FormValue("secondString")
	defer t.Execute(w, nil)
	fmt.Println("First:", firstString)
	fmt.Println("Second:", secondString)
}

func main() {
	http.HandleFunc("/", GetStrings)
	http.ListenAndServe(":8080", nil)

}
