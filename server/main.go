package main

import (
	"myapp/repositories"
	"myapp/db"
	"myapp/utils/validation"
	"myapp/utils/logic"
	"myapp/services"
	"myapp/controllers"
	"myapp/router"
)
func main() {
				//DBの初期化
				db := db.Init()
				//ロジックをインスタンス化
				jwtLogic := logic.NewJWTLogic()
				responseLogic := logic.NewResponseLogic()
				userLogic := logic.NewUserLogic()

				//バリデーションをインスタンス化
				authValidation := validation.NewAuthValidation()

				//リポジトリをインスタンス化
				userRepository := repositories.NewUserRepository(db)

				//サービスをインスタンス化
				authService := services.NewAuthService(responseLogic, userLogic, authValidation, userRepository, jwtLogic)

				//コントローラーをインスタンス化
				authController := controllers.NewAuthController(authService)

				//ルーティングをインスタンス化
				authRouter := router.NewAuthRouter(authController)
				mainRrouter := router.NewMainRouter(authRouter)
				mainRrouter.SetupRouting()
}