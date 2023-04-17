package controllers

import (
	"myapp/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController interface {
	SignUp(ctx echo.Context) error
}

type authController struct {
				as services.AuthService
}

func NewAuthController(as services.AuthService) AuthController {
				return &authController{as}
}

/*
	会員登録の処理
*/
func (authController *authController) SignUp(ctx echo.Context) error{
				w := ctx.Response().Writer
				r := ctx.Request()
				createUser, err := authController.as.SignUp(w, r)
				if err != nil {
					return err
				}
				authController.as.SendAuthResponse(w, &createUser, http.StatusCreated)
				return ctx.NoContent(http.StatusOK)
}