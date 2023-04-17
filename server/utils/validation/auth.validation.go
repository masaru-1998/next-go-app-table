package validation

import (
				"myapp/models"

				"github.com/go-ozzo/ozzo-validation/v4"
				"github.com/go-ozzo/ozzo-validation/v4/is"
)

type AuthValidation interface {
				SignUpParamValidation(param models.SignUpRequest ) error
}

type authValidation struct {}

func NewAuthValidation() AuthValidation{
				return &authValidation{}
}

/*
アカウント作成時のバリデーション
*/
func (av *authValidation) SignUpParamValidation(param models.SignUpRequest ) error {
				return validation.ValidateStruct(&param,
						validation.Field(&param.Name,
							validation.Required.Error("名前の入力は必須です"),
							validation.RuneLength(3, 24).Error("名前の入力は3~24文字です"),
						),
						validation.Field(&param.Email,
							validation.Required.Error("メールアドレスは必須です"),
							validation.RuneLength(8, 40).Error("メールアドレスの文字数は8~40です"),
							is.Email.Error("メールアドレスの形式が間違っています"),
						),
						validation.Field(&param.Password,
							validation.Required.Error("パスワードの入力は必須です"),
							validation.RuneLength(4, 8).Error("passwordは4~8文字です"),
							is.Alphanumeric.Error("英数字のみしか使えません"),
						),
				)
}