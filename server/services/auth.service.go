package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"myapp/models"
	"myapp/repositories"
	"myapp/utils/logic"
	"myapp/utils/validation"
	"net/http"

	"github.com/pkg/errors"
)

type AuthService interface {
					SignUp(w http.ResponseWriter, r *http.Request) (models.User, error)
					SendAuthResponse(w http.ResponseWriter, user *models.User, code int)
}

type authService struct {
				rl  logic.ResponseLogic
				ul  logic.UserLogic
				av  validation.AuthValidation
				ur  repositories.UserRepository
				jl  logic.JWTLogic
}

func NewAuthService(rl logic.ResponseLogic, ul logic.UserLogic, av validation.AuthValidation, ur repositories.UserRepository, jl logic.JWTLogic) AuthService{
	return &authService{rl, ul, av, ur, jl}
}

/*
 会員登録の処理
*/
func (as *authService) SignUp(w http.ResponseWriter, r *http.Request) (models.User, error) {
				body, _ := ioutil.ReadAll(r.Body)
				//bodyをjson形式から構造体に変更
				var signUpRequestParam models.SignUpRequest
				if err := json.Unmarshal(body, &signUpRequestParam); err != nil {
							log.Fatal(err)
							errorMessage := "リクエストボディに問題があるようです"
							as.rl.SendResponse(w, as.rl.CreateStringErrorResponse(errorMessage), http.StatusInternalServerError)
							return models.User{}, err
				}
				//バリデーションの処理
				if err := as.av.SignUpParamValidation(signUpRequestParam); err != nil {
					as.rl.SendResponse(w, as.rl.CreateErrorResponse(err), http.StatusBadRequest)
					return models.User{}, err
				}
				//ユーザがすでに存在しないかの確認
				var user []models.User
				if err := as.ur.GetAllUserByEmail(&user, signUpRequestParam.Email); err != nil {
								as.rl.SendResponse(w, as.rl.CreateErrorResponse(err), http.StatusInternalServerError)
								return  models.User{}, err
				}
				if len(user)!= 0 {
								as.rl.SendResponse(w, as.rl.CreateStringErrorResponse("入力されたメールアドレスはすでに存在しています。"), http.StatusUnauthorized)
								return models.User{}, errors.Errorf("「%w」 のユーザーは既に登録されています。", signUpRequestParam.Email)
				}
				//登録データの作成
				var createUser models.User
				hashPassword := as.ul.CreateHashPassword(signUpRequestParam.Password)
				createUser.Name = signUpRequestParam.Name
				createUser.Email = signUpRequestParam.Email
				createUser.Password = string(hashPassword)
				//ユーザ登録処理
				if err := as.ur.CreateUser(&createUser); err != nil {
							as.rl.SendResponse(w, as.rl.CreateStringErrorResponse("ユーザ登録に失敗しました"), http.StatusInternalServerError)
							return models.User{}, err
				}

				return createUser, nil
}

/*
会員登録API/ログインAPIのレスポンス送信処理
*/
func (as *authService) SendAuthResponse(w http.ResponseWriter, user *models.User, code int) {
				token, err := as.jl.CreateJwtToken(user)
				if err != nil {
					as.rl.SendResponse(w, as.rl.CreateStringErrorResponse("トークンの生成に失敗しました"), http.StatusInternalServerError)
					return
				}

				var response models.AuthResponse
				response.Token = token
				response.User.BaseModel.ID = user.ID
				response.User.BaseModel.CreatedAt = user.CreatedAt
				response.User.BaseModel.UpdatedAt = user.UpdatedAt
				response.User.BaseModel.DeletedAt = user.DeletedAt
				response.User.Name = user.Name
				response.User.Email = user.Email

				responseBody, _ := json.Marshal(response)
				as.rl.SendResponse(w, responseBody, code)
}