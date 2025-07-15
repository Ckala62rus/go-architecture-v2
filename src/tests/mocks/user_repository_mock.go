package mocks

import (
	"practice/domains"
)

// MockUserRepository is a mock implementation of the Users repository interface
type MockUserRepository struct {
	users []domains.User
}

// NewMockUserRepository creates a new mock user repository
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: []domains.User{},
	}
}

// SetUsers sets the users for the mock repository
func (m *MockUserRepository) SetUsers(users []domains.User) {
	m.users = users
}

// AddUser adds a user to the mock repository
func (m *MockUserRepository) AddUser(user domains.User) {
	m.users = append(m.users, user)
}

// GetUserByName implements the Users interface
func (m *MockUserRepository) GetUserByName(name string) (domains.User, error) {
	for _, user := range m.users {
		if user.Name == name {
			return user, nil
		}
	}
	return domains.User{}, NewNotFoundError("user with name: " + name + " not found")
}

// GetUserByEmail implements the Users interface
func (m *MockUserRepository) GetUserByEmail(email string) (domains.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return domains.User{}, NewNotFoundError("user with email: " + email + " not found")
}

// GetById implements the Users interface
func (m *MockUserRepository) GetById(id int) (domains.User, error) {
	for _, user := range m.users {
		if user.Id == id {
			return user, nil
		}
	}
	return domains.User{}, NewNotFoundError("user with id not found")
}

// GetAllUsers implements the Users interface
func (m *MockUserRepository) GetAllUsers() []domains.User {
	return m.users
}

// DeleteUserById implements the Users interface
func (m *MockUserRepository) DeleteUserById(id int) (bool, error) {
	for i, user := range m.users {
		if user.Id == id {
			// Remove user from slice
			m.users = append(m.users[:i], m.users[i+1:]...)
			return true, nil
		}
	}
	return false, NewNotFoundError("can't delete user with id")
}

// UpdateUser implements the Users interface
func (m *MockUserRepository) UpdateUser(userRequest domains.User) (domains.User, error) {
	for i, user := range m.users {
		if user.Id == userRequest.Id {
			// Update user data
			if userRequest.Name != "" {
				m.users[i].Name = userRequest.Name
			}
			if userRequest.Email != "" {
				m.users[i].Email = userRequest.Email
			}
			if userRequest.Password != "" {
				m.users[i].Password = userRequest.Password
			}
			return m.users[i], nil
		}
	}
	return domains.User{}, NewNotFoundError("user not found")
}

// Reset clears all users from the mock repository
func (m *MockUserRepository) Reset() {
	m.users = []domains.User{}
}
