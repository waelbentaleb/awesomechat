package stream

type Repository interface {
	InsertStream(record *Stream) error
	FindStream(username string) (*Stream, error)
	FindAllUsernames() ([]string, error)
	DeleteStream(username string) error
}
