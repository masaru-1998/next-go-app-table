package router

import (
	"github.com/labstack/echo/v4"

	"myapp/controllers"
)
type AuthRouter interface {
	SetAuthRouter(router *echo.Echo)
}

type authRouter struct {
			ac controllers.AuthController
}

func NewAuthRouter(ac controllers.AuthController) AuthRouter{
			return &authRouter{ac}
}

func (authRouter *authRouter) SetAuthRouter(router *echo.Echo) {
	/*
	body:
			name: string
			email: string
			password: string
	*/
			router.POST("/api/user/signup", echo.HandlerFunc(authRouter.ac.SignUp))
}
