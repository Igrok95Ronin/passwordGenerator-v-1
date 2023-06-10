package main

import (
	"log"
	"net/http"
	"text/template"
)

func passwordGenerator(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"../../ui/html/passwordGenerator.page.html",
		"../../ui/html/base.layout.html",
		"../../ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка на сервере", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Внутренняя ошибка на сервере2", 500)
		return
	}
}
