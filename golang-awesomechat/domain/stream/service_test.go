package stream_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/waelbentaleb/awesomechat/domain/group"
	"github.com/waelbentaleb/awesomechat/domain/stream"
	"github.com/waelbentaleb/awesomechat/domain/user"
	"github.com/waelbentaleb/awesomechat/storage/memory"
)

func TestService_Connect(t *testing.T) {
	userService, streamService, _ := initService()

	testUsername := "Wael"
	ctx := createNewUser(t, userService, testUsername)

	t.Run("OK", func(t *testing.T) {
		messageChannel, err := streamService.Connect(ctx, testUsername)
		assert.Nil(t, err)
		assert.NotNil(t, messageChannel)
	})

	t.Run("NOK - User already connected", func(t *testing.T) {
		messageChannel, err := streamService.Connect(ctx, testUsername)
		assert.NotNil(t, err)
		assert.Nil(t, messageChannel)
	})

	t.Run("NOK - Invalid username", func(t *testing.T) {
		messageChannel, err := streamService.Connect(ctx, "Wassim")
		assert.NotNil(t, err)
		assert.Nil(t, messageChannel)
	})
}

func TestService_DeleteStream(t *testing.T) {
	userService, streamService, _ := initService()

	testUsername := "Wael"
	ctx := createNewUser(t, userService, testUsername)

	t.Run("OK", func(t *testing.T) {
		_, err := streamService.Connect(ctx, testUsername)
		assert.Nil(t, err)

		err = streamService.DeleteStream(ctx, testUsername)
		assert.Nil(t, err)
	})

	t.Run("NOK - Stream already deleted", func(t *testing.T) {
		err := streamService.DeleteStream(ctx, testUsername)
		assert.Nil(t, err)
	})
}

func TestService_ListChannels(t *testing.T) {
	userService, streamService, _ := initService()

	testUsername := "Wael"
	ctx := createNewUser(t, userService, testUsername)

	t.Run("OK", func(t *testing.T) {
		_, err := streamService.Connect(ctx, testUsername)
		assert.Nil(t, err)

		users, groups, err := streamService.ListChannels(ctx)
		assert.Nil(t, err)
		assert.Equal(t, len(users), 1)
		assert.Equal(t, len(groups), 0)
	})
}

func TestService_SendMessage(t *testing.T) {
	userService, streamService, groupService := initService()

	testUsername1 := "Wael"
	ctx1 := createNewUser(t, userService, testUsername1)

	testUsername2 := "Wassim"
	ctx2 := createNewUser(t, userService, testUsername2)

	messageContent := "I'm a new message!"
	messageChannel := make(chan stream.Message)
	testGroupName := "Bingo"

	t.Run("OK - Direct message", func(t *testing.T) {
		_, err := streamService.Connect(ctx1, testUsername1)
		assert.Nil(t, err)

		messageChannel, err = streamService.Connect(ctx2, testUsername2)
		assert.Nil(t, err)

		go func() {
			msg := <-messageChannel
			log.Printf("Message from go routine: %s", msg.Content)
		}()

		err = streamService.SendMessage(ctx1, testUsername2, messageContent)
		assert.Nil(t, err)
	})

	t.Run("OK - Group message", func(t *testing.T) {
		err := groupService.CreateGroup(ctx1, testGroupName)
		assert.Nil(t, err)

		err = groupService.JoinGroup(ctx2, testGroupName)
		assert.Nil(t, err)

		go func() {
			msg := <-messageChannel
			log.Printf("Message from go routine: %s", msg.Content)
		}()

		err = streamService.SendMessage(ctx1, testGroupName, messageContent)
		assert.Nil(t, err)
	})

	t.Run("NOK - Invalid receiver", func(t *testing.T) {
		err := streamService.SendMessage(ctx1, "AAA", messageContent)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "receiver not exist")
	})
}

func initService() (*user.Service, *stream.Service, *group.Service) {
	userRepo := memory.NewUserRepository()
	streamRepo := memory.NewStreamRepository()
	groupRepo := memory.NewGroupRepository()

	userService := user.NewService(userRepo)
	streamService := stream.NewService(streamRepo, groupRepo)
	groupService := group.NewService(groupRepo, userRepo)

	return userService, streamService, groupService
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
