package group

type Repository interface {
	InsertGroup(record *Group) error
	FindGroup(groupName string) (*Group, error)
	FindAllGroupNames() ([]string, error)
	DeleteGroup(groupName string) error
	FindUserInGroup(groupName, username string) (bool, error)
	AddUserToGroup(groupName, username string) error
	RemoveUserFromGroup(groupName, username string) error
}
