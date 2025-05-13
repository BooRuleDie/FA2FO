package main

import (
	"fmt"
	"go-backend/internal/auth"
	"go-backend/internal/mailer"
	"go-backend/internal/store"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	activationURL := fmt.Sprintf("%s/confirm/%s", app.config.frontendURL, token)
	isProdEnv := app.config.env == "prod"
	vars := struct {
		Username      string
		ActivationURL string
	}{
		Username:      user.Username,
		ActivationURL: activationURL,
	}

	// send email
	err := app.mailer.Send(mailer.UserWelcomeTemplate, user.Username, user.Email, vars, !isProdEnv)
	if err != nil {
		app.logger.Errorw("error sending welcome email", "error", err)

		// rollback user creation if email fails (SAGA pattern)
		if err := app.store.Users.Delete(r.Context(), user.ID); err != nil {
			app.logger.Errorw("failed to delete user after welcome email fail", "error", err)
		}

		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, &uwt); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type CreateUserTokenPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

// createTokenHandler godoc
//
//	@Summary		Creates a token
//	@Description	Creates a token for a user
//	@Tags			authentication
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreateUserTokenPayload	true	"User credentials"
//	@Success		200		{string}	string					"Token"
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//
//	@Router			/auth/token [post]
func (app *application) createTokenHandler(w http.ResponseWriter, r *http.Request) {
	// parse payload
	var payload CreateUserTokenPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequest(w, r, err)
		return
	}

	// fetch the user data (check if user is existed)
	user, err := app.store.Users.GetByEmail(r.Context(), payload.Email)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.unauthorized(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}

		return
	}

	// generate the token
	claims := jwt.MapClaims{
		"sub": user.ID,
		"aud": auth.AUD,
		"exp": app.config.auth.token.exp,
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": auth.Hostname,
	}
	token, err := app.auth.GenerateToken(claims)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	// send it back to client
	if err := app.jsonResponse(w, http.StatusOK, token); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
