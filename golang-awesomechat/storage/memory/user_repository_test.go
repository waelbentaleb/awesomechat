package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/waelbentaleb/awesomechat/domain/user"
)

func TestUserRepository(t *testing.T) {
	userRepo := NewUserRepository()

	newUser := &user.User{
		Username: "Wael",
		Token:    "6694d2c422acd208a0072939487f6999eb9d18a44784045d",
	}

	t.Run("OK - Insert user", func(t *testing.T) {
		err := userRepo.InsertUser(newUser)
		assert.Nil(t, err)
	})

	t.Run("OK - Find user by username", func(t *testing.T) {
		res, err := userRepo.FindUserByUsername(newUser.Username)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, newUser.Token, res.Token)
	})

	t.Run("OK - Find user by token", func(t *testing.T) {
		res, err := userRepo.FindUserByToken(newUser.Token)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, newUser.Username, res.Username)
	})

	t.Run("OK - Delete user", func(t *testing.T) {
		err := userRepo.DeleteUser(newUser.Username)
		assert.Nil(t, err)

		res, err := userRepo.FindUserByUsername(newUser.Username)
		assert.Nil(t, err)
		assert.Nil(t, res)
	})
}
