package jwt

import (
	"time"

	"github.com/Brolervi/Twitter_Chafa/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenJWT, Generates a JSON WEB TOKEN
func GenJWT(t models.User) (string, error) {
	pass := []byte("TimeNeverWaits_ItDeliversAllEquallyToTheSameEnd_YouWhoWishToSafeguardTheFutureHoweverLimitedItMayBe_YouWillBeGivenOneYear_GoForthWithoutFalterWithYourHeartAsYourGuide") // Secret Pass

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lname":     t.LName,
		"birthdate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(pass)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
