package main

import (
	"log"
	"net/http"
)

const appPort = ":8080"

type application struct {
	addr string
}

func (s *application) getUsersHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("user data fetched!"))
}

func (s *application) createUsersHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("user created!"))
}

func main() {
	app := &application{addr: appPort}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", app.getUsersHandler)
	mux.HandleFunc("POST /users", app.createUsersHandler)

	srv := http.Server{
		Addr:    app.addr,
		Handler: mux,
	}
	log.Println("server started listening!")
	log.Fatal(srv.ListenAndServe())
}
