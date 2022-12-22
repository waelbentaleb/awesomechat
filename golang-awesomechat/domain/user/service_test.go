package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/waelbentaleb/awesomechat/domain/user"
	"github.com/waelbentaleb/awesomechat/storage/memory"
)

func TestService_CreateUser(t *testing.T) {
	service := newUserService()
	ctx := context.Background()
	testUsername := "Wael"

	t.Run("OK", func(t *testing.T) {
		token, err := service.CreateUser(ctx, testUsername)
		assert.Nil(t, err)
		assert.NotNil(t, token)
	})

	t.Run("NOK - Invalid username", func(t *testing.T) {
		token, err := service.CreateUser(ctx, "")

		assert.Nil(t, token)
		assert.Equal(t, "please enter a valid username", err.Error())
	})

	t.Run("NOK - username exist", func(t *testing.T) {
		token, err := service.CreateUser(ctx, testUsername)
		assert.NotNil(t, err)
		assert.Nil(t, token)
	})

}

func TestService_ValidateToken(t *testing.T) {
	service := newUserService()
	ctx := context.Background()
	testUsername := "Wael"
	testToken := "6694d2c422acd208a0072939487f6999eb9d18a44784045d"

	t.Run("OK", func(t *testing.T) {
		token, err := service.CreateUser(ctx, testUsername)
		assert.Nil(t, err)
		assert.NotNil(t, token)

		testUser, err := service.ValidateToken(*token)
		assert.Nil(t, err)
		assert.NotNil(t, testUser)
		assert.Equal(t, testUser.Username, testUsername)
	})

	t.Run("NOK - Invalid token", func(t *testing.T) {
		testUser, err := service.ValidateToken(testToken)
		assert.NotNil(t, err)
		assert.Nil(t, testUser)
	})
}

func newUserService() *user.Service {
	userRepo := memory.NewUserRepository()
	userService := user.NewService(userRepo)

	return userService
}
