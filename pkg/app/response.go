package app

import (
	"github.com/go-chi/render"
	"net/http"
	"trueconf-userapi/models"
)

const (
	SUCCESS        = "SUCCESS"
	UPDATE_SUCCESS = "success updating"
	DELETE_SUCCESS = "success deleting"

	UNKNOWN_ERROR           = "unknown error"
	ERROR_BAD_URL_PARAMETER = "bad url parameter"
	ERROR_MISSING_PARAMETER = "parameter missed"

	ERROR_GET_USER         = "user does not exists"
	ERROR_BAD_REQUEST_TYPE = "bad type of request"
	ERROR_USER_NOT_FOUND   = models.USER_NOT_FOUND

	MESSAGE_WELCOME = "welcome to trueconf-userapi. Try /api/v1/users"

	CODE_SUCCESS      = 200
	CODE_CLIENT_ERROR = 400
	CODE_SERVER_ERROR = 500
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(w http.ResponseWriter, r *http.Request, code int, message string, data interface{}) {
	render.Status(r, code)
	render.JSON(w, r, Response{Code: code, Message: message, Data: data})
}
