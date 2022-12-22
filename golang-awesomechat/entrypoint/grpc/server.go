package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/waelbentaleb/awesomechat/contracts"

	"github.com/waelbentaleb/awesomechat/domain/group"
	"github.com/waelbentaleb/awesomechat/domain/stream"
	"github.com/waelbentaleb/awesomechat/domain/user"
)

type AwesomeChatServer struct {
	userService   *user.Service
	streamService *stream.Service
	groupService  *group.Service

	pb.UnimplementedChatCoreServer
}

func NewAwesomeChatServer(
	userService *user.Service,
	streamService *stream.Service,
	groupService *group.Service,
) *AwesomeChatServer {
	return &AwesomeChatServer{
		userService:   userService,
		streamService: streamService,
		groupService:  groupService,
	}
}

func (s *AwesomeChatServer) CreateUser(ctx context.Context, user *pb.User) (*pb.CreateUserResponse, error) {

	token, err := s.userService.CreateUser(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{Token: *token}, nil
}

func (s *AwesomeChatServer) Connect(user *pb.User, stream pb.ChatCore_ConnectServer) error {
	ch, err := s.streamService.Connect(stream.Context(), user.Username)
	if err != nil {
		return err
	}

	for message := range ch {
		err := stream.Send(&pb.ReceivedMessage{
			Sender:    message.Sender,
			Content:   message.Content,
			Type:      message.Type,
			Groupname: message.GroupName,
			Date:      timestamppb.New(message.Date),
		})

		if err != nil {
			return err
		}
	}

	err = s.streamService.DeleteStream(stream.Context(), user.Username)
	if err != nil {
		return err
	}

	return nil
}

func (s *AwesomeChatServer) SendMessage(ctx context.Context, message *pb.SentMessage) (*pb.Empty, error) {

	err := s.streamService.SendMessage(ctx, message.Receiver, message.Content)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *AwesomeChatServer) CreateGroupChat(ctx context.Context, req *pb.Group) (*pb.Empty, error) {

	err := s.groupService.CreateGroup(ctx, req.Groupname)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *AwesomeChatServer) JoinGroupChat(ctx context.Context, req *pb.Group) (*pb.Empty, error) {

	err := s.groupService.JoinGroup(ctx, req.Groupname)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *AwesomeChatServer) LeftGroupChat(ctx context.Context, req *pb.Group) (*pb.Empty, error) {

	err := s.groupService.LeftGroup(ctx, req.Groupname)
	if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *AwesomeChatServer) ListChannels(ctx context.Context, req *pb.Empty) (*pb.ListChannelsResponse, error) {

	users, groups, err := s.streamService.ListChannels(ctx)
	if err != nil {
		return nil, err
	}

	var items []*pb.ListChannelsItem

	for _, u := range users {
		items = append(items, &pb.ListChannelsItem{
			Type:       "USER",
			Identifier: u,
		})
	}

	for _, g := range groups {
		items = append(items, &pb.ListChannelsItem{
			Type:       "GROUP",
			Identifier: g,
		})
	}

	return &pb.ListChannelsResponse{Items: items}, nil
}
