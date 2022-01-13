package app

import (
	"net/http"
	"trueconf-userapi/models"
)

type UserRequest models.User

func (c *UserRequest) Bind(r *http.Request) error { return nil }
