// service_test.go
package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract/repository"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/repository/mocks"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestUserService_Get(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := service.NewUserService(&repository.Contract{User: mockRepo})

	// mocking data
	mockUsers := []model.User{
		{ID: uuid.New(), Name: "John Doe", Email: "johndoe@me.com"},
		{ID: uuid.New(), Name: "Jane Doe", Email: "janedoe@me.com"},
	}

	mockRepo.On("Get", mock.Anything).Return(&mockUsers, nil)
	result, err := userService.Get(context.Background())

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, &mockUsers, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetById_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := service.NewUserService(&repository.Contract{User: mockRepo})

	// moching data
	mockUser := &model.User{ID: uuid.New(), Name: "John Doe"}

	mockRepo.On("GetById", mock.Anything, mockUser.ID).Return(mockUser, nil)
	result, err := userService.GetById(context.Background(), mockUser.ID)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, mockUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetById_NotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := service.NewUserService(&repository.Contract{User: mockRepo})

	// Mengatur mock agar mengembalikan error
	mockRepo.On("GetById", mock.Anything, mock.Anything).Return((*model.User)(nil), gorm.ErrRecordNotFound)
	result, err := userService.GetById(context.Background(), uuid.New())

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Create_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := service.NewUserService(&repository.Contract{User: mockRepo})

	mockUser := &model.User{Name: "Alice", Password: "password123"}

	// Mengatur mock agar mengembalikan ID simulasi
	mockRepo.On("Create", mock.Anything, mock.Anything).Return("new_id", nil)
	id, err := userService.Create(context.Background(), mockUser)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "new_id", id)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Update(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := service.NewUserService(&repository.Contract{User: mockRepo})

	// mocking data
	id := uuid.New()
	payload := &model.User{
		ID:       id,
		Name:     "Updated Name",
		Email:    "updated@example.com",
		Password: "newpassword",
	}

	// Mengatur mock untuk mengembalikan data yang sudah ada
	mockRepo.On("GetById", mock.Anything, id).Return(&model.User{ID: id}, nil)
	mockRepo.On("Update", mock.Anything, payload, id).Return("Updated Successfully", nil)
	result, err := userService.Update(context.Background(), payload, id)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, "Updated Successfully", result)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Delete(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userService := service.NewUserService(&repository.Contract{User: mockRepo})

	// ID yang akan dihapus
	id := uuid.New()

	// Mengatur mock untuk memanggil Delete tanpa error
	mockRepo.On("Delete", mock.Anything, id).Return(nil)

	// Memanggil service untuk menghapus data
	err := userService.Delete(context.Background(), id)

	// AserAssertionssi
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
