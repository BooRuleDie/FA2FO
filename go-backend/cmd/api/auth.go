package main

import (
	"crypto/sha256"
	"encoding/hex"
	"go-backend/internal/store"
	"net/http"

	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

// registerUserHandler godoc
//
//	@Summary		Registers a user
//	@Description	Registers a user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		RegisterUserPayload	true	"User credentials"
//	@Success		201		{object}	store.User			“User	registered"
//	@Failure		400		{object}	error

// @Failure	500	{object}	error
// @Router		/authentication/user [post]
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
	plainToken := uuid.New().String()
	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])

	if err := app.store.Users.CreateAndInvite(r.Context(), user, hashToken, app.config.mail.exp); err != nil {
		// handle errors
		if err != nil {
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
	}

	if err := app.jsonResponse(w, http.StatusOK, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
