package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/waelbentaleb/awesomechat/contracts"

	"github.com/waelbentaleb/awesomechat/domain/group"
	"github.com/waelbentaleb/awesomechat/domain/stream"
	"github.com/waelbentaleb/awesomechat/domain/user"
	grpcserver "github.com/waelbentaleb/awesomechat/entrypoint/grpc"
	"github.com/waelbentaleb/awesomechat/middleware"
	"github.com/waelbentaleb/awesomechat/storage/memory"
)

func main() {
	fmt.Println("Hello Awesome Chat ^_^ ")

	userRepository := memory.NewUserRepository()
	streamRepository := memory.NewStreamRepository()
	groupRepository := memory.NewGroupRepository()

	userService := user.NewService(userRepository)
	streamService := stream.NewService(streamRepository, groupRepository)
	groupService := group.NewService(groupRepository, userRepository)

	authInterceptor := middleware.NewAuthInterceptor(userService)
	server := grpc.NewServer(
		authInterceptor.UnaryAuthInterceptor(),
		authInterceptor.StreamAuthInterceptor(),
	)

	awesomeChatServer := grpcserver.NewAwesomeChatServer(userService, streamService, groupService)
	pb.RegisterChatCoreServer(server, awesomeChatServer)

	// Register reflection service on gRPC server.
	reflection.Register(server)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
