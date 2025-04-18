package main

import (
	"go-backend/internal/store"
	"net/http"

	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}

// registerUserHandler godoc
//
//	@Summary		Registers a user
//	@Description	Registers a user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterUserPayload	true	"User credentials"
//	@Success		201		{object}	UserWithToken		"User	registered"
//	@Failure		400		{object}	error
//	@Failure		409		{object}	error	"Duplicate email or username"
//	@Failure		500		{object}	error
//	@Router			/auth/user [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var rup RegisterUserPayload
	if err := readJSON(w, r, &rup); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := Validate.Struct(rup); err != nil {
		app.badRequest(w, r, err)
		return
	}

	// hash the user password
	user := &store.User{
		Username: rup.Username,
		Email:    rup.Email,
	}

	if err := user.Password.Set(rup.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// generate token
	token := uuid.New().String()

	if err := app.store.Users.CreateAndInvite(r.Context(), user, token, app.config.mail.exp); err != nil {
		// handle errors
		switch err {
		case store.ErrDuplicateEmail:
			app.badRequest(w, r, store.ErrDuplicateEmail)
			return
		case store.ErrDuplicateUsername:
			app.badRequest(w, r, store.ErrDuplicateUsername)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	uwt := UserWithToken{
		User:  user,
		Token: token,
	}

	if err := app.jsonResponse(w, http.StatusOK, &uwt); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
