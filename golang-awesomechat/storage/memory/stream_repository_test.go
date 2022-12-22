package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/waelbentaleb/awesomechat/domain/stream"
	"github.com/waelbentaleb/awesomechat/domain/user"
)

func TestStreamRepository(t *testing.T) {
	streamRepo := NewStreamRepository()

	userOne := &user.User{
		Username: "Wael",
		Token:    "6694d2c422acd208a0072939487f6999eb9d18a44784045d",
	}

	userTwo := &user.User{
		Username: "Wassim",
		Token:    "7bbb0407d1e2c64981855ad8681d0d86d1e91e00167939cb",
	}

	streamOne := &stream.Stream{
		Username:       "Wael",
		MessageChannel: make(chan stream.Message),
	}

	streamTwo := &stream.Stream{
		Username:       "Wassim",
		MessageChannel: make(chan stream.Message),
	}

	t.Run("OK - Insert stream", func(t *testing.T) {
		err := streamRepo.InsertStream(streamOne)
		assert.Nil(t, err)

		err = streamRepo.InsertStream(streamTwo)
		assert.Nil(t, err)
	})

	t.Run("OK - Find stream", func(t *testing.T) {
		res, err := streamRepo.FindStream(userOne.Username)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, res.MessageChannel, streamOne.MessageChannel)
	})

	t.Run("OK - Delete stream", func(t *testing.T) {
		err := streamRepo.DeleteStream(userOne.Username)
		assert.Nil(t, err)

		res, err := streamRepo.FindStream(userOne.Username)
		assert.Nil(t, err)
		assert.Nil(t, res)

		res, err = streamRepo.FindStream(userTwo.Username)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("OK - Find all streams", func(t *testing.T) {
		res, err := streamRepo.FindAllUsernames()
		assert.Nil(t, err)
		assert.Equal(t, len(res), 1)
	})
}
