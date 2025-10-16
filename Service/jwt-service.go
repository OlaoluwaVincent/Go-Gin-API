package service

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(userID string, isAdmin bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJwtService() JwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "my_app",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (jstSrv *jwtService) GenerateToken(username string, isAdmin bool) string {
	claims := &jwtCustomClaims{
		username,
		isAdmin,
		jwt.StandardClaims{
			Issuer: jstSrv.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jstSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (jstSrv *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jstSrv.secretKey), nil
	})
}
