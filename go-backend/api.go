package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

const appPort = ":8080"

type application struct {
	addr string
}

var users = make([]user, 0)

func (s *application) getUsersHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	
	if err := json.NewEncoder(rw).Encode(users); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		// this line is actually quite important
		// just by forgetting return in those cases can cuase
		// critical security vulnerabilities
		return 
	}
	
	rw.WriteHeader(http.StatusOK)
}

func (s *application) createUsersHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var newUser user
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	if err := insert(&newUser); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return 
	}
	
	rw.WriteHeader(http.StatusCreated)
}

func insert(u *user) error {
	if u.Firstname == "" || u.Lastname == ""  {
		return errors.New("firstname and lastname fields can't be empty")
	}
	
	for _, savedUser := range users {
		if savedUser.Firstname == u.Firstname && savedUser.Lastname == u.Lastname {
			return errors.New("user already exists")
		}
	}
	
	users = append(users, *u)
	
	return nil
}