package stream

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/waelbentaleb/awesomechat/domain/group"
	"github.com/waelbentaleb/awesomechat/domain/user"
)

const (
	DirectMessage = "DIRECT"
	GroupMessage  = "GROUP"
)

type Service struct {
	streamRepo Repository
	groupRepo  group.Repository
}

func NewService(
	streamRepo Repository,
	groupRepo group.Repository,
) *Service {
	return &Service{
		streamRepo: streamRepo,
		groupRepo:  groupRepo,
	}
}

// Connect takes a username and create a new messages channel used to stream messages to connected user
func (s *Service) Connect(cxt context.Context, username string) (chan Message, error) {

	// Check that the provided token is associated with the given username
	currentUser := cxt.Value("user").(*user.User)
	if currentUser.Username != username {
		return nil, errors.New("invalid username")
	}

	stream, err := s.streamRepo.FindStream(username)
	if err != nil {
		return nil, err
	}

	// Check if a message channel is already created
	if stream != nil {
		return nil, errors.New("user already connected")
	}

	// Create new stream object with message channel and store it
	messageChannel := make(chan Message)
	stream = &Stream{
		Username:       username,
		MessageChannel: messageChannel,
	}

	err = s.streamRepo.InsertStream(stream)
	if err != nil {
		return nil, err
	}

	log.Printf("%s connect to awesome chat stream", username)

	return messageChannel, nil
}

// SendMessage is used to send message to a user or a group of users
func (s *Service) SendMessage(cxt context.Context, receiver string, msg string) error {
	if msg == "" {
		return errors.New("please enter a valid message")
	}

	currentUser := cxt.Value("user").(*user.User)
	if receiver == currentUser.Username {
		return errors.New("you can't send a message to your self")
	}

	// Check if the receiver is a registered user
	stream, err := s.streamRepo.FindStream(receiver)
	if err != nil && err.Error() != "stream not found" {
		return err
	}

	// Send a direct message in the case when a stream exist for a given username
	if stream != nil {
		messageChannel := stream.MessageChannel

		message := Message{
			Content: msg,
			Sender:  currentUser.Username,
			Type:    DirectMessage,
			Date:    time.Now(),
		}

		messageChannel <- message
		log.Printf("%s send direct message to %s", currentUser.Username, receiver)

		return nil
	}

	// Check if the receiver is a registered group channel
	groupChat, err := s.groupRepo.FindGroup(receiver)
	if err != nil {
		return err
	}

	// Return an error when no group found with the given receiver name
	if groupChat == nil {
		// if the receiver is not a registered user and not a registered group channel we return an error
		return errors.New("receiver not exist")
	}

	// Check that the current user is joined the target group
	inGroup, err := s.groupRepo.FindUserInGroup(groupChat.GroupName, currentUser.Username)
	if err != nil {
		return err
	}
	if !inGroup {
		return errors.New("you are not a member of the group")
	}

	// Prepare group message
	message := Message{
		Content:   msg,
		Sender:    currentUser.Username,
		Type:      GroupMessage,
		GroupName: groupChat.GroupName,
		Date:      time.Now(),
	}

	// Loop into group members and send the message to all joined users except the sender
	for _, receiverUsername := range groupChat.JoinedUsers {
		if receiverUsername == currentUser.Username {
			continue
		}

		userStream, err := s.streamRepo.FindStream(receiverUsername)
		if err != nil {
			log.Printf("error sending message from %s to %s", currentUser.Username, receiverUsername)
		}

		userStream.MessageChannel <- message
		log.Printf("%s send a group message to %s", currentUser.Username, receiverUsername)
	}

	return nil
}

// DeleteStream takes a username and destroy the associated steam
func (s *Service) DeleteStream(cxt context.Context, username string) error {

	currentUser := cxt.Value("user").(*user.User)
	if currentUser.Username != username {
		return errors.New("not authorized")
	}

	err := s.streamRepo.DeleteStream(username)
	if err != nil {
		return err
	}

	return nil
}

// ListChannels list all connected users and available groups
func (s *Service) ListChannels(cxt context.Context) ([]string, []string, error) {
	users, err := s.streamRepo.FindAllUsernames()
	if err != nil {
		return nil, nil, err
	}

	groups, err := s.groupRepo.FindAllGroupNames()
	if err != nil {
		return nil, nil, err
	}

	return users, groups, nil
}
