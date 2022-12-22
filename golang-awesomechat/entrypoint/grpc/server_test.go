package grpc_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/waelbentaleb/awesomechat/contracts"

	"github.com/waelbentaleb/awesomechat/domain/group"
	"github.com/waelbentaleb/awesomechat/domain/stream"
	"github.com/waelbentaleb/awesomechat/domain/user"
	"github.com/waelbentaleb/awesomechat/entrypoint/grpc"
	"github.com/waelbentaleb/awesomechat/storage/memory"
)

func TestAwesomeChatServer_CreateUser(t *testing.T) {
	userService, streamService, groupService := initService()
	ctx := context.Background()

	testUsername := "Wael"
	testUser := &pb.User{
		Username: testUsername,
	}

	t.Run("OK", func(t *testing.T) {
		server := grpc.NewAwesomeChatServer(userService, streamService, groupService)
		resp, err := server.CreateUser(ctx, testUser)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestAwesomeChatServer_SendMessage(t *testing.T) {
	userService, streamService, groupService := initService()

	testUsername1 := "Wael"
	ctx1 := createNewUser(t, userService, testUsername1)

	testUsername2 := "Wassim"
	ctx2 := createNewUser(t, userService, testUsername2)

	messageContent := &pb.SentMessage{
		Content:  "Hello friend!",
		Receiver: testUsername2,
	}

	t.Run("OK", func(t *testing.T) {
		messageChannel, err := streamService.Connect(ctx2, testUsername2)
		assert.Nil(t, err)

		go func() {
			msg := <-messageChannel
			log.Printf("Message from go routine: %s", msg.Content)
		}()

		server := grpc.NewAwesomeChatServer(userService, streamService, groupService)
		resp, err := server.SendMessage(ctx1, messageContent)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestAwesomeChatServer_CreateGroupChat(t *testing.T) {
	userService, streamService, groupService := initService()

	testUsername := "Wael"
	ctx := createNewUser(t, userService, testUsername)

	testGroupName := "Bingo"
	testGroup := &pb.Group{
		Groupname: testGroupName,
	}

	t.Run("OK", func(t *testing.T) {
		server := grpc.NewAwesomeChatServer(userService, streamService, groupService)
		resp, err := server.CreateGroupChat(ctx, testGroup)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestAwesomeChatServer_JoinGroupChat(t *testing.T) {
	userService, streamService, groupService := initService()

	testUsername1 := "Wael"
	ctx1 := createNewUser(t, userService, testUsername1)

	testUsername2 := "Wassim"
	ctx2 := createNewUser(t, userService, testUsername2)

	testGroupName := "Bingo"
	testGroup := &pb.Group{
		Groupname: testGroupName,
	}

	t.Run("OK", func(t *testing.T) {
		err := groupService.CreateGroup(ctx1, testGroupName)
		assert.Nil(t, err)

		server := grpc.NewAwesomeChatServer(userService, streamService, groupService)
		resp, err := server.JoinGroupChat(ctx2, testGroup)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestAwesomeChatServer_LeftGroupChat(t *testing.T) {
	userService, streamService, groupService := initService()

	testUsername1 := "Wael"
	ctx1 := createNewUser(t, userService, testUsername1)

	testGroupName := "Bingo"
	testGroup := &pb.Group{
		Groupname: testGroupName,
	}

	t.Run("OK", func(t *testing.T) {
		err := groupService.CreateGroup(ctx1, testGroupName)
		assert.Nil(t, err)

		server := grpc.NewAwesomeChatServer(userService, streamService, groupService)
		resp, err := server.LeftGroupChat(ctx1, testGroup)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestAwesomeChatServer_ListChannels(t *testing.T) {
	userService, streamService, groupService := initService()

	testUsername := "Wael"
	ctx := createNewUser(t, userService, testUsername)

	testGroupName := "Bingo"
	empty := &pb.Empty{}

	t.Run("OK", func(t *testing.T) {
		_, err := streamService.Connect(ctx, testUsername)
		assert.Nil(t, err)

		err = groupService.CreateGroup(ctx, testGroupName)
		assert.Nil(t, err)

		server := grpc.NewAwesomeChatServer(userService, streamService, groupService)
		resp, err := server.ListChannels(ctx, empty)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 2, len(resp.Items))
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
