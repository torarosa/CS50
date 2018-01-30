package main

import (
	"net/http"
	"templating"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var g []string
	if r.Method == http.MethodPost {
		firstString := r.FormValue("firstString")
		secondString := r.FormValue("secondString")
		g = []string{firstString, secondString}
	}

	templating.RenderTemplate(w, "index", g)
}

func compare(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	templating.RenderTemplate(w, "compare", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/compare", compare)
	http.ListenAndServe(":8080", nil)
}
