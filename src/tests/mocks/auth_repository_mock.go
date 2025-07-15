package mocks

import (
	"practice/domains"
)

// MockAuthRepository is a mock implementation of the Authorization repository interface
type MockAuthRepository struct {
	users  []domains.User
	nextId int
}

// NewMockAuthRepository creates a new mock auth repository
func NewMockAuthRepository() *MockAuthRepository {
	return &MockAuthRepository{
		users:  []domains.User{},
		nextId: 1,
	}
}

// SetUsers sets the users for the mock repository
func (m *MockAuthRepository) SetUsers(users []domains.User) {
	m.users = users
	// Update nextId to be max(user.Id) + 1
	maxId := 0
	for _, user := range users {
		if user.Id > maxId {
			maxId = user.Id
		}
	}
	m.nextId = maxId + 1
}

// CreateUser implements the Authorization interface
func (m *MockAuthRepository) CreateUser(user domains.User) (domains.User, error) {
	// Check if user already exists by email
	for _, existingUser := range m.users {
		if existingUser.Email == user.Email {
			return domains.User{}, NewValidationError("email", "user already exists")
		}
	}

	// Assign ID and add to users
	user.Id = m.nextId
	m.nextId++
	m.users = append(m.users, user)

	return user, nil
}

// GetUserByEmail is a helper method for testing
func (m *MockAuthRepository) GetUserByEmail(email string) (domains.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return domains.User{}, NewNotFoundError("user with email: " + email + " not found")
}

// Reset clears all users from the mock repository
func (m *MockAuthRepository) Reset() {
	m.users = []domains.User{}
	m.nextId = 1
}

// GetAllUsers returns all users (for testing purposes)
func (m *MockAuthRepository) GetAllUsers() []domains.User {
	return m.users
}
