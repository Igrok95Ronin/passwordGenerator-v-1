package main

import (
	"log"
	"net/http"
	"text/template"
)

// Home
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"../../ui/html/home.page.html",
		"../../ui/html/baseHome.layout.html",
		"../../ui/html/footer.partial.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренная ошибка на сервере", 500)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка на сервере2", 500)
		return
	}
}
