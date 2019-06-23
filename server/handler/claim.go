package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ymohl-cl/gopkg/server/model"
)

// user calims key
const (
	claimUserIDKey     = "identifier"
	claimUserPseudoKey = "pseudo"
	claimUserAgeKey    = "age"
	claimUserGenreKey  = "genre"
	claimUserEmailKey  = "email"
)

func (h Handler) generateToken(user model.User) (string, error) {
	var token *jwt.Token
	var hash string
	var err error

	// create token
	token = jwt.New(jwt.SigningMethodHS256)
	// sets claims
	claims := token.Claims.(jwt.MapClaims)
	claims[claimUserIDKey] = user.Identifier
	claims[claimUserPseudoKey] = user.Pseudo
	claims[claimUserAgeKey] = user.Age
	claims[claimUserGenreKey] = user.Genre
	claims[claimUserEmailKey] = user.Email
	if hash, err = token.SignedString([]byte(h.JWTKey)); err != nil {
		return "", err
	}
	return hash, nil
}
