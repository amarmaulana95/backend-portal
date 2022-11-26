package auth

import (
	"fmt"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

type Service interface { //membuat interfaces service
	GenerateToken(userID int) (string, error) //generate token paramnya userID
}

type jwtService struct {
}

var SCREET_KEY = []byte("B15mill4h_K3y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) { //membuat fungsi Generate token, balikannya string dan error

	Expired := time.Now().Add(time.Hour * 24 * 7).Unix()
	fmt.Println(Expired)

	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = Expired

	//claim  jwt

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SCREET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil

}
