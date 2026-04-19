package service

import (
	"fmt"

	"github.com/aigerim25/Goo/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"testing"
)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Ilimessova Aigerim"}
	mockRepo.EXPECT().GetUserById(1).Return(user, nil)

	result, err := userService.GetUserById(1)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}
func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Ilimessova Aigerim"}
	mockRepo.EXPECT().CreateUser(user).Return(nil)
	err := userService.CreateUser(user)

	assert.NoError(t, err)
}

// Test cases for RegisterUser
func TestRegisterUserWithExistingEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Ilimessova Aigerim"}
	email := "aigerim@gmail.com"

	mockRepo.EXPECT().GetByEmail(email).Return(user, nil)
	err := userService.RegisterUser(user, email)
	assert.Error(t, err)
	assert.Equal(t, "user with this email already exists", err.Error())
}
func TestNewUserSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Ilimessova Aigerim"}
	email := "aigerim@gmail.com"

	mockRepo.EXPECT().GetByEmail(email).Return(nil, nil)
	mockRepo.EXPECT().CreateUser(user).Return(nil)

	err := userService.RegisterUser(user, email)
	assert.NoError(t, err)
}
func TestRepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Ilimessova Aigerim"}
	email := "aigerim@gmail.com"

	mockRepo.EXPECT().GetByEmail(email).Return(nil, fmt.Errorf("error"))
	err := userService.RegisterUser(user, email)

	assert.Error(t, err)
	assert.Equal(t, "error getting user with this email", err.Error())
}

// Test cases for UpdateUserName
func TestCheckingForEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	err := userService.UpdateUserName(1, "")

	assert.Error(t, err)
	assert.Equal(t, "new name cannot be empty", err.Error())
}
func TestUserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	mockRepo.EXPECT().GetUserById(1).Return(nil, fmt.Errorf("user not found"))
	err := userService.UpdateUserName(1, "Aiko")

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}
func SuccessfulUpdateUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Ilimessova Aigerim"}
	mockRepo.EXPECT().GetUserById(1).Return(user, nil)
	mockRepo.EXPECT().CreateUser(user).Return(nil)

	err := userService.UpdateUserName(1, "Aiko")

	assert.NoError(t, err)
	assert.Equal(t, "Aiko", user.Name)
}
func TestUpdateUserNameFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{
		ID:   1,
		Name: "Ilimessova Aigerim",
	}
	mockRepo.EXPECT().GetUserById(1).Return(user, nil)
	mockRepo.EXPECT().UpdateUser(user).Return(fmt.Errorf("update failed"))

	err := userService.UpdateUserName(1, "Aiko")
	assert.Error(t, err)
	assert.Equal(t, "update failed", err.Error())
}
func TestChangesBeforeUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Ilimessova Aigerim"}

	mockRepo.EXPECT().GetUserById(1).Return(user, nil)
	mockRepo.EXPECT().UpdateUser(gomock.Any()).DoAndReturn(func(updatedUser *repository.User) error {
		assert.Equal(t, "Aiko", updatedUser.Name)
		return nil
	})
	err := userService.UpdateUserName(1, "Aiko")
	assert.NoError(t, err)
}
