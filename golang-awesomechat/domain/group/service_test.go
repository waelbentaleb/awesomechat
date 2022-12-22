package group_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/waelbentaleb/awesomechat/domain/group"
	"github.com/waelbentaleb/awesomechat/domain/user"
	"github.com/waelbentaleb/awesomechat/storage/memory"
)

func TestService_CreateGroup(t *testing.T) {
	userService, groupService := initService()
	testGroupName := "Bingo"

	testUsername := "Wael"
	ctx := createNewUser(t, userService, testUsername)

	t.Run("OK", func(t *testing.T) {
		err := groupService.CreateGroup(ctx, testGroupName)

		assert.Nil(t, err)
	})

	t.Run("NOK - Invalid group name", func(t *testing.T) {
		err := groupService.CreateGroup(ctx, "")

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid group name")
	})

	t.Run("NOK - Group name exist", func(t *testing.T) {
		err := groupService.CreateGroup(ctx, testGroupName)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "group name already exist")
	})
}

func TestService_JoinGroup(t *testing.T) {
	userService, groupService := initService()
	testGroupName := "Bingo"

	testUsername1 := "Wael"
	ctx1 := createNewUser(t, userService, testUsername1)

	testUsername2 := "Wassim"
	ctx2 := createNewUser(t, userService, testUsername2)

	t.Run("OK", func(t *testing.T) {
		err := groupService.CreateGroup(ctx1, testGroupName)
		assert.Nil(t, err)

		err = groupService.JoinGroup(ctx2, testGroupName)
		assert.Nil(t, err)
	})

	t.Run("NOK - Already joined", func(t *testing.T) {
		err := groupService.JoinGroup(ctx1, testGroupName)

		assert.Equal(t, err.Error(), "already joined the group")
	})

	t.Run("NOK - Group dont exist", func(t *testing.T) {
		err := groupService.JoinGroup(ctx1, "AAA")
		assert.Equal(t, err.Error(), "group dont exist")
	})
}

func TestService_LeftGroup(t *testing.T) {
	userService, groupService := initService()
	testGroupName := "Bingo"

	testUsername1 := "Wael"
	ctx1 := createNewUser(t, userService, testUsername1)

	testUsername2 := "Wassim"
	ctx2 := createNewUser(t, userService, testUsername2)

	t.Run("OK", func(t *testing.T) {
		err := groupService.CreateGroup(ctx1, testGroupName)
		assert.Nil(t, err)

		err = groupService.LeftGroup(ctx1, testGroupName)
		assert.Nil(t, err)
	})

	t.Run("NOK - Group dont exist", func(t *testing.T) {
		err := groupService.CreateGroup(ctx1, testGroupName)
		assert.Nil(t, err)

		err = groupService.LeftGroup(ctx1, "AAA")
		assert.NotNil(t, err)
	})

	t.Run("NOK - Must join the group", func(t *testing.T) {
		err := groupService.LeftGroup(ctx2, testGroupName)
		assert.NotNil(t, err)
	})
}

func initService() (*user.Service, *group.Service) {
	userRepo := memory.NewUserRepository()
	groupRepo := memory.NewGroupRepository()

	userService := user.NewService(userRepo)
	groupService := group.NewService(groupRepo, userRepo)

	return userService, groupService
}

func createNewUser(t *testing.T, userService *user.Service, username string) context.Context {
	ctx := context.Background()
	token, err := userService.CreateUser(ctx, username)
	assert.Nil(t, err)

	ctx = context.WithValue(ctx, "user", &user.User{
		Username: username,
		Token:    *token,
	})

	return ctx
}
