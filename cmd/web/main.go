package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP") // Создаем новый флаг командной строки, значение по умолчанию: ":4000"
	flag.Parse()                                               // Мы вызываем функцию flag.Parse() для извлечения флага из командной строки

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	//Calculator
	mux.HandleFunc("/calculator", calculator)
	mux.HandleFunc("/formHandler", formHandler)
	mux.HandleFunc("/deleteEntry", deleteEntry)
	//ToDoList
	mux.HandleFunc("/todolist", toDoList)
	mux.HandleFunc("/formHandlerToDoList", formHandlerToDoList)
	mux.HandleFunc("/deleteEntryToDoList", deleteEntryToDoList)
	mux.HandleFunc("/editPost", editPost)

	//Password Generator
	mux.HandleFunc("/passwordgenerator", passwordGenerator)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("../../ui/static/")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Запуск сервера на %s", *addr) // Значение, возвращаемое функцией flag.String(), является указателем на значение go run ./cmd/web -addr=":9999"
	err := http.ListenAndServe(*addr, mux)
	log.Fatalln(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
