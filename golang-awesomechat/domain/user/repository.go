package user

type Repository interface {
	InsertUser(user *User) error
	DeleteUser(username string) error
	FindUserByUsername(username string) (*User, error)
	FindUserByToken(token string) (*User, error)
}
