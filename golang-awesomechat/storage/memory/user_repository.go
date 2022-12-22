package memory

import (
	"sync"

	"github.com/waelbentaleb/awesomechat/domain/user"
)

// This is a representation of in memory users store
// Users are stored in a map of "token" as Key and "username" as Value

type UserRepository struct {
	users map[string]string
	mu    sync.Mutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]string),
	}
}

func (r *UserRepository) InsertUser(record *user.User) error {
	r.mu.Lock()
	r.users[record.Token] = record.Username
	r.mu.Unlock()

	return nil
}

func (r *UserRepository) FindUserByUsername(username string) (*user.User, error) {
	r.mu.Lock()
	for token, uname := range r.users {
		if uname == username {
			r.mu.Unlock()
			return &user.User{
				Username: username,
				Token:    token,
			}, nil
		}
	}

	r.mu.Unlock()
	return nil, nil
}

func (r *UserRepository) FindUserByToken(token string) (*user.User, error) {
	r.mu.Lock()
	username, ok := r.users[token]
	r.mu.Unlock()

	if !ok {
		return nil, nil
	}

	return &user.User{
		Username: username,
		Token:    token,
	}, nil
}

func (r *UserRepository) DeleteUser(username string) error {
	r.mu.Lock()
	for token, uname := range r.users {
		if uname == username {
			delete(r.users, token)
			r.mu.Unlock()
			return nil
		}
	}

	r.mu.Unlock()
	return nil
}
