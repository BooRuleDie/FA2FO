package main

import (
	"log"
	"net/http"
)

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
