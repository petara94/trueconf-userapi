package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"trueconf-userapi/config"
)

func TestUserStore_Create(t *testing.T) {
	st := UserStore{
		StorePath: "../" + config.UserStoreFile,
	}

	user := User{
		DisplayName: "Peter",
		Email:       "mail@ru",
	}

	_, err := st.Create(user)
	assert.Nil(t, err)
	err = st.Delete(user.ID)
	assert.Nil(t, err)
}

func TestUserStore_Delete(t *testing.T) {
	st := UserStore{
		StorePath: "../" + config.UserStoreFile,
	}

	anyIdReturnNil := 1111

	// Ошибка возникает лишь при повреждении/отсутствии файла бд
	err := st.Delete(anyIdReturnNil)
	assert.Nil(t, err)
}

func TestUserStore_Update(t *testing.T) {
	st := UserStore{
		StorePath: "../" + config.UserStoreFile,
	}

	user := User{
		DisplayName: "Peter",
		Email:       "mail@ru",
	}

	user, err := st.Create(user)
	assert.Nil(t, err)

	newUser := User{
		DisplayName: "Peter2",
		Email:       "gmail@com",
	}

	err = st.Update(user.ID, newUser)
	assert.Nil(t, err)
}

func TestUserStore_Get(t *testing.T) {
	st := UserStore{
		StorePath: "../" + config.UserStoreFile,
	}

	user := User{
		DisplayName: "Peter",
		Email:       "mail@ru",
	}

	user, err := st.Create(user)
	assert.Nil(t, err)

	user2, err := st.Get(user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user, user2)
}

func TestUserStore_GetFail(t *testing.T) {
	st := UserStore{
		StorePath: "../" + config.UserStoreFile,
	}

	_, err := st.Get(-9999)
	assert.NotNil(t, err)
}
