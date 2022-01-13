package user_service

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"trueconf-userapi/config"
	"trueconf-userapi/models"
	"trueconf-userapi/pkg/app"
)

type UserService struct {
	store models.UserStore
}

func NewUserService() *UserService {
	return &UserService{store: *models.NewUserStore(config.UserStoreFile)}
}

func (us *UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	UserIdString := chi.URLParam(r, "id")
	if UserIdString == "" {
		app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_MISSING_PARAMETER, nil)
		return
	}
	UserId, err := strconv.Atoi(UserIdString)
	if err != nil {
		app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_BAD_URL_PARAMETER, nil)
		return
	}

	user, err := us.store.Get(UserId)
	if err != nil {
		app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_GET_USER, nil)
		return
	}

	app.SendResponse(w, r, app.CODE_SUCCESS, app.SUCCESS, user)
}

func (us *UserService) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := us.store.GetAll()
	if users == nil {
		app.SendResponse(w, r, app.CODE_SERVER_ERROR, app.UNKNOWN_ERROR, nil)
		return
	}

	app.SendResponse(w, r, app.CODE_SUCCESS, app.SUCCESS, users)
}

func (us *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := app.UserRequest{}

	err := render.Bind(r, &newUser)
	if err != nil {
		app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_BAD_REQUEST_TYPE, nil)
		return
	}

	user, err := us.store.Create(models.User(newUser))

	if err != nil {
		app.SendResponse(w, r, app.CODE_SERVER_ERROR, app.UNKNOWN_ERROR, nil)
		return
	}

	app.SendResponse(w, r, app.CODE_SUCCESS, app.SUCCESS, user)
}

func (us *UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	UserIdString := chi.URLParam(r, "id")
	if UserIdString == "" {
		app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_MISSING_PARAMETER, nil)
		return
	}
	UserId, err := strconv.Atoi(UserIdString)
	if err != nil {
		app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_BAD_URL_PARAMETER, nil)
		return
	}

	updUser := app.UserRequest{}

	err = render.Bind(r, &updUser)
	if err != nil {
		app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_BAD_REQUEST_TYPE, nil)
		return
	}

	err = us.store.Update(UserId, models.User(updUser))
	if err != nil {
		if err.Error() == app.ERROR_USER_NOT_FOUND {
			app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_USER_NOT_FOUND, nil)
		} else {
			app.SendResponse(w, r, app.CODE_SERVER_ERROR, app.UNKNOWN_ERROR, nil)
		}
		return
	}

	app.SendResponse(w, r, app.CODE_SUCCESS, app.UPDATE_SUCCESS, nil)
}

func (us *UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	UserIdString := chi.URLParam(r, "id")
	if UserIdString == "" {
		app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_MISSING_PARAMETER, nil)
		return
	}
	UserId, err := strconv.Atoi(UserIdString)
	if err != nil {
		app.SendResponse(w, r, app.CODE_CLIENT_ERROR, app.ERROR_BAD_URL_PARAMETER, nil)
		return
	}

	err = us.store.Delete(UserId)
	if err != nil {
		app.SendResponse(w, r, app.CODE_SERVER_ERROR, app.UNKNOWN_ERROR, nil)
		return
	}

	app.SendResponse(w, r, app.CODE_SUCCESS, app.DELETE_SUCCESS, nil)
}
