package group

import (
	"context"
	"errors"
	"log"

	"github.com/waelbentaleb/awesomechat/domain/user"
)

type Service struct {
	groupRepo Repository
	userRepo  user.Repository
}

func NewService(
	groupRepo Repository,
	userRepo user.Repository,
) *Service {
	return &Service{
		groupRepo: groupRepo,
		userRepo:  userRepo,
	}
}

// CreateGroup takes a group name and create a new group with the current user as first member
func (s *Service) CreateGroup(cxt context.Context, groupName string) error {
	if groupName == "" {
		return errors.New("invalid group name")
	}

	// Check if there is a registered user with the given group name
	// group should not have the same name as stored users usernames
	existUser, err := s.userRepo.FindUserByUsername(groupName)
	if err != nil {
		return err
	}
	if existUser != nil {
		return errors.New("invalid group name")
	}

	existGroup, err := s.groupRepo.FindGroup(groupName)
	if err != nil {
		return err
	}
	if existGroup != nil {
		return errors.New("group name already exist")
	}

	// Create a new group with current user as first member
	currentUser := cxt.Value("user").(*user.User)
	newGroup := &Group{
		GroupName:   groupName,
		JoinedUsers: []string{currentUser.Username},
	}

	err = s.groupRepo.InsertGroup(newGroup)
	if err != nil {
		return err
	}

	log.Printf("%s group created by %s", groupName, currentUser.Username)

	return nil
}

// JoinGroup allow users to join an exist group by providing a group name
func (s *Service) JoinGroup(cxt context.Context, groupName string) error {
	group, err := s.groupRepo.FindGroup(groupName)
	if err != nil {
		return err
	}
	if group == nil {
		return errors.New("group dont exist")
	}

	currentUser := cxt.Value("user").(*user.User)
	inGroup, err := s.groupRepo.FindUserInGroup(groupName, currentUser.Username)
	if err != nil {
		return err
	}
	if inGroup {
		return errors.New("already joined the group")
	}

	err = s.groupRepo.AddUserToGroup(groupName, currentUser.Username)
	if err != nil {
		return err
	}

	log.Printf("%s joined %s group", currentUser.Username, groupName)

	return nil
}

// LeftGroup allow users to leave a group by providing a group name
func (s *Service) LeftGroup(cxt context.Context, groupName string) error {
	group, err := s.groupRepo.FindGroup(groupName)
	if err != nil {
		return err
	}
	if group == nil {
		return errors.New("group dont exist")
	}

	currentUser := cxt.Value("user").(*user.User)
	inGroup, err := s.groupRepo.FindUserInGroup(groupName, currentUser.Username)
	if err != nil {
		return err
	}
	if !inGroup {
		return errors.New("you are not joined the group")
	}

	err = s.groupRepo.RemoveUserFromGroup(groupName, currentUser.Username)

	if len(group.JoinedUsers) == 1 {
		log.Printf("Group %s will be deleted", group.GroupName)
		err = s.groupRepo.DeleteGroup(groupName)
	}

	log.Printf("%s left %s group", currentUser.Username, groupName)

	return nil
}
