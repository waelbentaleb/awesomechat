package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waelbentaleb/awesomechat/domain/group"
)

func TestGroupRepository(t *testing.T) {
	groupRepo := NewGroupRepository()

	testGroupOne := &group.Group{
		GroupName:   "Bingo",
		JoinedUsers: []string{"Wael", "Wassim"},
	}

	testGroupTwo := &group.Group{
		GroupName:   "Chella",
		JoinedUsers: []string{"Wael", "Salim", "Firas"},
	}

	t.Run("OK - Insert group", func(t *testing.T) {
		err := groupRepo.InsertGroup(testGroupOne)
		assert.Nil(t, err)

		err = groupRepo.InsertGroup(testGroupTwo)
		assert.Nil(t, err)
	})

	t.Run("OK - Find group by group name", func(t *testing.T) {
		res, err := groupRepo.FindGroup(testGroupOne.GroupName)
		assert.Nil(t, err)
		assert.Equal(t, res.GroupName, testGroupOne.GroupName)
		assert.Equal(t, len(res.JoinedUsers), 2)
	})

	t.Run("OK - Find all group names", func(t *testing.T) {
		res, err := groupRepo.FindAllGroupNames()
		assert.Nil(t, err)
		assert.Equal(t, len(res), 2)
	})

	t.Run("OK - Delete group", func(t *testing.T) {
		err := groupRepo.DeleteGroup(testGroupOne.GroupName)
		assert.Nil(t, err)

		res, err := groupRepo.FindAllGroupNames()
		assert.Nil(t, err)
		assert.Equal(t, len(res), 1)
	})

	t.Run("OK - Find group member", func(t *testing.T) {
		res, err := groupRepo.FindUserInGroup(testGroupTwo.GroupName, "Firas")
		assert.Nil(t, err)
		assert.Equal(t, res, true)
	})

	t.Run("OK - Add user to group", func(t *testing.T) {
		err := groupRepo.AddUserToGroup(testGroupTwo.GroupName, "Salah")
		assert.Nil(t, err)

		res, err := groupRepo.FindUserInGroup(testGroupTwo.GroupName, "Salah")
		assert.Nil(t, err)
		assert.Equal(t, res, true)
	})

	t.Run("OK - Remove user from group", func(t *testing.T) {
		err := groupRepo.RemoveUserFromGroup(testGroupTwo.GroupName, "Wael")
		assert.Nil(t, err)

		res, err := groupRepo.FindUserInGroup(testGroupTwo.GroupName, "Wael")
		assert.Nil(t, err)
		assert.Equal(t, res, false)
	})
}
