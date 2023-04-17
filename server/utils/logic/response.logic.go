package logic

import (
	"encoding/json"
	"net/http"
)

type ResponseLogic interface {
				SendResponse(w http.ResponseWriter, response []byte, code int)
				CreateStringErrorResponse( errorMessage string) []byte
				CreateErrorResponse(err error) []byte
}

type responseLogic struct {}

func NewResponseLogic() ResponseLogic {
				return &responseLogic{}
}

/*
レスポンス内容の設定
*/
func (responseLogic *responseLogic) SendResponse(w http.ResponseWriter, response []byte, code int) {
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(code)
				w.Write(response)
}

/*
エラーメッセージを生成
*/
func (responseLogic *responseLogic) CreateErrorResponse(err error) []byte {
				response := map[string]interface{}{
					"error": err,
				}
				responseBody, _ := json.Marshal(response)

				return responseBody
}

/*
エラーメッセージを整形
*/
func (responseLogic *responseLogic) CreateStringErrorResponse( errorMessage string) []byte {
				response := map[string]interface{}{
								"error": errorMessage,
				}
				responseBody, _ := json.Marshal(response)
				return responseBody
}

