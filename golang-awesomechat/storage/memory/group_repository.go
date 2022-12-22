package memory

import (
	"errors"
	"sync"

	"github.com/waelbentaleb/awesomechat/domain/group"
)

// This is a representation of in memory group store
// Groups are stored in a map of "group name" as Key and array of "joined users" as Value

type GroupRepository struct {
	groups map[string][]string
	mu     sync.Mutex
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{
		groups: make(map[string][]string),
	}
}

func (r *GroupRepository) InsertGroup(record *group.Group) error {
	r.mu.Lock()
	r.groups[record.GroupName] = record.JoinedUsers
	r.mu.Unlock()

	return nil
}

func (r *GroupRepository) FindGroup(groupName string) (*group.Group, error) {
	r.mu.Lock()
	joinedUsers, ok := r.groups[groupName]
	r.mu.Unlock()

	if !ok {
		return nil, nil
	}

	return &group.Group{
		GroupName:   groupName,
		JoinedUsers: joinedUsers,
	}, nil
}

func (r *GroupRepository) FindAllGroupNames() ([]string, error) {
	var groupList []string

	r.mu.Lock()
	for groupChat, _ := range r.groups {
		groupList = append(groupList, groupChat)
	}
	r.mu.Unlock()

	return groupList, nil
}

func (r *GroupRepository) DeleteGroup(groupName string) error {
	r.mu.Lock()
	delete(r.groups, groupName)
	r.mu.Unlock()

	return nil
}

func (r *GroupRepository) FindUserInGroup(groupName, username string) (bool, error) {
	r.mu.Lock()
	joinedUsers, ok := r.groups[groupName]
	r.mu.Unlock()

	if !ok {
		return false, errors.New("group not found")
	}

	for _, user := range joinedUsers {
		if user == username {
			return true, nil
		}
	}

	return false, nil
}

func (r *GroupRepository) AddUserToGroup(groupName, username string) error {
	r.mu.Lock()
	_, ok := r.groups[groupName]

	if !ok {
		return errors.New("group not found")
	}

	r.groups[groupName] = append(r.groups[groupName], username)
	r.mu.Unlock()

	return nil
}

func (r *GroupRepository) RemoveUserFromGroup(groupName, username string) error {
	r.mu.Lock()
	joinedUsers, ok := r.groups[groupName]

	if !ok {
		return errors.New("group not found")
	}

	index := -1

	for i, user := range joinedUsers {
		if user == username {
			index = i
			break
		}
	}

	if index < 0 {
		return errors.New("user not found")
	}

	joinedUsers[index] = joinedUsers[len(joinedUsers)-1]
	r.groups[groupName] = joinedUsers[:len(joinedUsers)-1]

	r.mu.Unlock()
	return nil
}
