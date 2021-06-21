package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	err = godotenv.Load()
	key = os.Getenv("SECRET_KEY")
)

type service interface {
	GenerateToken(UserID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{
		"user_id": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
