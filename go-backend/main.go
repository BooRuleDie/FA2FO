package main

import (
	"log"
	"net/http"
)

const appPort = ":8080"

type application struct {
	addr string
}

func (s *application) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			rw.Write([]byte("index page"))
			return
		case "/users":
			rw.Write([]byte("users endpoint"))
			return
		default:
			rw.Write([]byte("unknown path"))
			return
		}
	default:
		rw.Write([]byte("method not allowed"))
		return
	}

}

func main() {
	app := &application{addr: appPort}
	srv := &http.Server{
		Addr: app.addr,
		Handler: app,
	}
	
	log.Println("server started listening!")
	log.Fatal(srv.ListenAndServe())
}
