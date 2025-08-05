package services

import (
	"practice/domains"
	"practice/tests/mocks"
	"testing"
	"time"
)

// UserServiceInterface для тестирования
type UserServiceInterface interface {
	GetUserByName(name string) (domains.User, error)
	GetById(id int) (domains.User, error)
	GetAllUsers() []domains.User
	DeleteUserById(id int) (bool, error)
	UpdateUser(userRequest domains.User) (domains.User, error)
}

// SimpleUserService - упрощенная реализация для тестирования
type SimpleUserService struct {
	repo *mocks.MockUserRepository
}

func NewSimpleUserService(repo *mocks.MockUserRepository) *SimpleUserService {
	return &SimpleUserService{repo: repo}
}

func (u *SimpleUserService) GetUserByName(name string) (domains.User, error) {
	return u.repo.GetUserByName(name)
}

func (u *SimpleUserService) GetById(id int) (domains.User, error) {
	return u.repo.GetById(id)
}

func (u *SimpleUserService) GetAllUsers() []domains.User {
	return u.repo.GetAllUsers()
}

func (u *SimpleUserService) DeleteUserById(id int) (bool, error) {
	return u.repo.DeleteUserById(id)
}

func (u *SimpleUserService) UpdateUser(userRequest domains.User) (domains.User, error) {
	return u.repo.UpdateUser(userRequest)
}

// Тесты для SimpleUserService
func TestSimpleUserService_GetUserByName(t *testing.T) {
	// Setup
	mockRepo := mocks.NewMockUserRepository()
	service := NewSimpleUserService(mockRepo)

	// Test data
	testUser := domains.User{
		Id:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.AddUser(testUser)

	// Test cases
	tests := []struct {
		name         string
		userName     string
		expectedUser domains.User
		expectedErr  bool
	}{
		{
			name:         "existing user",
			userName:     "John Doe",
			expectedUser: testUser,
			expectedErr:  false,
		},
		{
			name:         "non-existing user",
			userName:     "Non Existing",
			expectedUser: domains.User{},
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetUserByName(tt.userName)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result.Id != tt.expectedUser.Id {
				t.Errorf("Expected Id %d, got %d", tt.expectedUser.Id, result.Id)
			}
			if result.Name != tt.expectedUser.Name {
				t.Errorf("Expected Name %s, got %s", tt.expectedUser.Name, result.Name)
			}
			if result.Email != tt.expectedUser.Email {
				t.Errorf("Expected Email %s, got %s", tt.expectedUser.Email, result.Email)
			}
		})
	}
}

func TestSimpleUserService_GetById(t *testing.T) {
	// Setup
	mockRepo := mocks.NewMockUserRepository()
	service := NewSimpleUserService(mockRepo)

	// Test data
	testUser := domains.User{
		Id:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.AddUser(testUser)

	// Test cases
	tests := []struct {
		name         string
		userId       int
		expectedUser domains.User
		expectedErr  bool
	}{
		{
			name:         "existing user",
			userId:       1,
			expectedUser: testUser,
			expectedErr:  false,
		},
		{
			name:         "non-existing user",
			userId:       999,
			expectedUser: domains.User{},
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetById(tt.userId)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result.Id != tt.expectedUser.Id {
				t.Errorf("Expected Id %d, got %d", tt.expectedUser.Id, result.Id)
			}
		})
	}
}

func TestSimpleUserService_GetAllUsers(t *testing.T) {
	// Setup
	mockRepo := mocks.NewMockUserRepository()
	service := NewSimpleUserService(mockRepo)

	// Test data
	testUsers := []domains.User{
		{
			Id:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "hashedpassword1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        2,
			Name:      "Jane Smith",
			Email:     "jane@example.com",
			Password:  "hashedpassword2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepo.SetUsers(testUsers)

	// Test
	result := service.GetAllUsers()

	if len(result) != len(testUsers) {
		t.Errorf("Expected %d users, got %d", len(testUsers), len(result))
	}

	for i, user := range result {
		if user.Id != testUsers[i].Id {
			t.Errorf("Expected user Id %d, got %d", testUsers[i].Id, user.Id)
		}
		if user.Name != testUsers[i].Name {
			t.Errorf("Expected user Name %s, got %s", testUsers[i].Name, user.Name)
		}
	}
}

func TestSimpleUserService_DeleteUserById(t *testing.T) {
	// Setup
	mockRepo := mocks.NewMockUserRepository()
	service := NewSimpleUserService(mockRepo)

	// Test data
	testUser := domains.User{
		Id:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.AddUser(testUser)

	// Test cases
	tests := []struct {
		name        string
		userId      int
		expectedRes bool
		expectedErr bool
	}{
		{
			name:        "existing user",
			userId:      1,
			expectedRes: true,
			expectedErr: false,
		},
		{
			name:        "non-existing user",
			userId:      999,
			expectedRes: false,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.DeleteUserById(tt.userId)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result != tt.expectedRes {
				t.Errorf("Expected result %v, got %v", tt.expectedRes, result)
			}
		})
	}
}

func TestSimpleUserService_UpdateUser(t *testing.T) {
	// Setup
	mockRepo := mocks.NewMockUserRepository()
	service := NewSimpleUserService(mockRepo)

	// Test data
	testUser := domains.User{
		Id:        1,
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "hashedpassword",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.AddUser(testUser)

	// Test cases
	tests := []struct {
		name          string
		updateRequest domains.User
		expectedName  string
		expectedErr   bool
	}{
		{
			name: "update existing user",
			updateRequest: domains.User{
				Id:   1,
				Name: "John Updated",
			},
			expectedName: "John Updated",
			expectedErr:  false,
		},
		{
			name: "update non-existing user",
			updateRequest: domains.User{
				Id:   999,
				Name: "Non Existing",
			},
			expectedName: "",
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.UpdateUser(tt.updateRequest)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result.Name != tt.expectedName {
				t.Errorf("Expected Name %s, got %s", tt.expectedName, result.Name)
			}
		})
	}
}
