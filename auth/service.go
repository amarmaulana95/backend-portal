package auth

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/dgrijalva/jwt-go.v3"
)

type Service interface { //membuat interfaces service
	GenerateToken(userID int) (string, error)       //generate token paramnya userID
	ValidateToken(token string) (*jwt.Token, error) //untuk validasi token
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

func (s *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) { // validasi token
	// masukan token , parameternya adalah func lalu mengembalikan interface dan err
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		//cek tokennya
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SCREET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil

}
