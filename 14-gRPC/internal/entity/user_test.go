package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	name := "Stefan"
	email := "XU4yN@example.com"
	password := "password"
	user, err := NewUser(name, email, password)
	require.Nil(t, err)
	require.NotNil(t, user)
	require.Equal(t, name, user.Name)
	require.Equal(t, email, user.Email)
	require.NotEmpty(t, user.Password)
	require.NotEmpty(t, user.ID)
}

func TestUser_ValidatePassword(t *testing.T) {
	name := "Stefan"
	email := "XU4yN@example.com"
	password := "123456"
	user, err := NewUser(name, email, password)
	require.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)

}
