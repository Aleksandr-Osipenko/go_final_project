package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

const (
	webDir = "./web" //дирекрория веб-приложения
)

func main() {
   //определение порта для запуска приложения
   port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

   //
	r := chi.NewRouter()
   r.Handle("/*", http.FileServer(http.Dir(webDir)))
	//r.Get("/js/scripts.min.js", handleScripts)
	

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}
