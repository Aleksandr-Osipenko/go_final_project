package main

import (
	"net/http"
	

	"github.com/go-chi/chi/v5"
)

const (
	webDir = "./web"
//	port   = "7540"
)

func main() {

	r := chi.NewRouter()
   r.Handle("/*", http.FileServer(http.Dir(webDir)))
	//r.Get("/js/scripts.min.js", handleScripts)
	

	err := http.ListenAndServe(":7540", r)
	if err != nil {
		panic(err)
	}
}
