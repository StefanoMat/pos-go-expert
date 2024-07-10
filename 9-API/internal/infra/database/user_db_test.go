package database

import (
	"testing"

	"github.com/stefanomat/pos-go-expert/9-API/internal/entity"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	require.NoError(t, err)
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Stefan", "XU4yN@example.com", "password")
	userDB := NewUser(db)
	err = userDB.Create(user)
	require.Nil(t, err)
	require.NotEmpty(t, user.ID)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	require.Nil(t, err)

	require.Equal(t, user.ID, userFound.ID)
	require.Equal(t, user.Name, userFound.Name)
	require.Equal(t, user.Email, userFound.Email)
	require.NotNil(t, userFound.Password)

}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"))
	require.NoError(t, err)
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Stefan", "XU4yN@example.com", "password")
	userDB := NewUser(db)
	err = userDB.Create(user)
	require.Nil(t, err)
	require.NotEmpty(t, user.ID)

	userFound, err := userDB.FindByEmail("XU4yN@example.com")
	require.Nil(t, err)
	require.NotEmpty(t, userFound.ID)
	require.Equal(t, user.ID, userFound.ID)
	require.Equal(t, user.Name, userFound.Name)
	require.Equal(t, user.Email, userFound.Email)
	require.NotNil(t, userFound.Password)
}
