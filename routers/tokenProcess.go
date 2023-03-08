package routers

import (
	"errors"
	"strings"

	"github.com/Brolervi/Twitter_Chafa/db"
	"github.com/Brolervi/Twitter_Chafa/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email, email value used in all endpoints
var Email string

// IDUser, returned ID of the model, that will be used in all endpoints
var IDUser string

// TokenProcess, extracts token values
func TokenProcess(tk string) (*models.Claim, bool, string, error) {
	pass := []byte("PilotoAviadorNoMuereSiNoVuelaMasAlto")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return pass, nil
	})
	if err == nil {
		_, found, _ := db.CheckUserAlreadyExists(claims.Email)
		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
