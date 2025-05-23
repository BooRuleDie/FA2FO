package auth

import "github.com/golang-jwt/jwt/v5"

var (
	AUD     string = "GopherSocial"
	Hostname string = "GopherSocial"
)

type JWTAuthenticator struct {
	secret string
	aud    string
	iss    string
}

func NewJWTAuthenticator(secret, aud, iss string) *JWTAuthenticator {
	return &JWTAuthenticator{
		secret: secret,
		aud:    aud,
		iss:    iss,
	}
}

func (auth *JWTAuthenticator) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(auth.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (auth *JWTAuthenticator) ValidateToken(token string) (*jwt.Token, error) {
	return nil, nil
}
