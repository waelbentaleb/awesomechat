package memory

import (
	"sync"

	"github.com/waelbentaleb/awesomechat/domain/stream"
)

// This is a representation of in memory stream store
// Streams are stored in a map of "username" as Key and "channel of messages" as Value

type StreamRepository struct {
	streams map[string]chan stream.Message
	mu      sync.Mutex
}

func NewStreamRepository() *StreamRepository {
	return &StreamRepository{
		streams: make(map[string]chan stream.Message),
	}
}

func (r *StreamRepository) InsertStream(record *stream.Stream) error {
	r.mu.Lock()
	r.streams[record.Username] = record.MessageChannel
	r.mu.Unlock()

	return nil
}

func (r *StreamRepository) FindStream(username string) (*stream.Stream, error) {
	r.mu.Lock()
	channel, ok := r.streams[username]
	r.mu.Unlock()

	if !ok {
		return nil, nil
	}

	return &stream.Stream{
		Username:       username,
		MessageChannel: channel,
	}, nil
}

func (r *StreamRepository) FindAllUsernames() ([]string, error) {
	var userList []string

	r.mu.Lock()
	for username, _ := range r.streams {
		userList = append(userList, username)
	}
	r.mu.Unlock()

	return userList, nil
}

func (r *StreamRepository) DeleteStream(username string) error {
	r.mu.Lock()
	delete(r.streams, username)
	r.mu.Unlock()

	return nil
}
