package logic

import (
	"fmt"
	"myapp/models"
	"os"
	"strconv"
	"time"

	// jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/joho/godotenv"
)

type JWTLogic interface {
				CreateJwtToken(user *models.User) (string, error)
}

type jwtLogic struct {}

func NewJWTLogic() JWTLogic{
				return &jwtLogic{}
}

func (jwtLogic *jwtLogic) CreateJwtToken(user *models.User) (string, error) {
				token := jwt.New(jwt.SigningMethodHS256)//トークンの雛形作成
				claims :=  token.Claims.(jwt.MapClaims)//claimの型アサーション
				claims["admin"] = true
				claims["sub"]   = strconv.Itoa(int(user.ID)) + user.Email + user.Name
				claims["name"]  = user.Name
				claims["exp"]   = time.Now().Add(time.Hour * 24).Unix()

				err := godotenv.Load()
				if err != nil {
								fmt.Println(err)
								return "", err
				}
				tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_KEY")))//設定した暗号化法を用いて文字列化
				return tokenString, nil
}