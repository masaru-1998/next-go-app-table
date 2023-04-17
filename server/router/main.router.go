package router

import(
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MainRouter interface{
				SetupRouting()
}

type mainRouter struct{
				authR AuthRouter
}

/*
*コンストラクタ
*/
func NewMainRouter(authR AuthRouter) MainRouter{
	return &mainRouter{authR}
}

/*
*ルーティングの定義
*/
func (mainRouter *mainRouter) SetupRouting(){
				router := echo.New()
				router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
					AllowOrigins:     []string{"http://localhost:3000"},
					AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
					AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
					AllowCredentials: true,
				}))
				mainRouter.authR.SetAuthRouter(router)
				fmt.Println("API Server started")
				router.Logger.Fatal(router.Start(":8080"))
}