package dto

import (
	"practice/domains"
	"practice/pkg/dto"
	"testing"
	"time"
)

func TestMapSingleUser(t *testing.T) {
	// Prepare test data
	testTime := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	testUser := domains.User{
		Id:        1,
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Password:  "hashedPassword123",
		CreatedAt: testTime,
		UpdatedAt: testTime,
	}

	expected := dto.UserOutDTO{
		Id:        1,
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		CreatedAt: "2024-01-15 10:30:00",
		UpdatedAt: "2024-01-15 10:30:00",
	}

	// Test the mapping
	result := dto.MapSingleUser(testUser)

	// Verify results
	if result.Id != expected.Id {
		t.Errorf("MapSingleUser() Id = %v, want %v", result.Id, expected.Id)
	}
	if result.Name != expected.Name {
		t.Errorf("MapSingleUser() Name = %v, want %v", result.Name, expected.Name)
	}
	if result.Email != expected.Email {
		t.Errorf("MapSingleUser() Email = %v, want %v", result.Email, expected.Email)
	}
	if result.CreatedAt != expected.CreatedAt {
		t.Errorf("MapSingleUser() CreatedAt = %v, want %v", result.CreatedAt, expected.CreatedAt)
	}
	if result.UpdatedAt != expected.UpdatedAt {
		t.Errorf("MapSingleUser() UpdatedAt = %v, want %v", result.UpdatedAt, expected.UpdatedAt)
	}
}

func TestMapSingleUserWithEmptyValues(t *testing.T) {
	// Test with empty/zero values
	testUser := domains.User{
		Id:        0,
		Name:      "",
		Email:     "",
		Password:  "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	result := dto.MapSingleUser(testUser)

	if result.Id != 0 {
		t.Errorf("MapSingleUser() Id = %v, want %v", result.Id, 0)
	}
	if result.Name != "" {
		t.Errorf("MapSingleUser() Name = %v, want %v", result.Name, "")
	}
	if result.Email != "" {
		t.Errorf("MapSingleUser() Email = %v, want %v", result.Email, "")
	}

	// Check that dates are formatted correctly even when zero
	expectedDate := "0001-01-01 00:00:00"
	if result.CreatedAt != expectedDate {
		t.Errorf("MapSingleUser() CreatedAt = %v, want %v", result.CreatedAt, expectedDate)
	}
	if result.UpdatedAt != expectedDate {
		t.Errorf("MapSingleUser() UpdatedAt = %v, want %v", result.UpdatedAt, expectedDate)
	}
}

func TestMapAllUser(t *testing.T) {
	// Prepare test data
	testTime1 := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	testTime2 := time.Date(2024, 1, 16, 11, 45, 0, 0, time.UTC)

	testUsers := []domains.User{
		{
			Id:        1,
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			Password:  "hashedPassword123",
			CreatedAt: testTime1,
			UpdatedAt: testTime1,
		},
		{
			Id:        2,
			Name:      "Jane Smith",
			Email:     "jane.smith@example.com",
			Password:  "hashedPassword456",
			CreatedAt: testTime2,
			UpdatedAt: testTime2,
		},
	}

	expected := []dto.UserOutDTO{
		{
			Id:        1,
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: "2024-01-15 10:30:00",
			UpdatedAt: "2024-01-15 10:30:00",
		},
		{
			Id:        2,
			Name:      "Jane Smith",
			Email:     "jane.smith@example.com",
			CreatedAt: "2024-01-16 11:45:00",
			UpdatedAt: "2024-01-16 11:45:00",
		},
	}

	// Test the mapping
	result := dto.MapAllUser(testUsers)

	// Verify results
	if len(result) != len(expected) {
		t.Errorf("MapAllUser() length = %v, want %v", len(result), len(expected))
		return
	}

	for i := range result {
		if result[i].Id != expected[i].Id {
			t.Errorf("MapAllUser() [%d] Id = %v, want %v", i, result[i].Id, expected[i].Id)
		}
		if result[i].Name != expected[i].Name {
			t.Errorf("MapAllUser() [%d] Name = %v, want %v", i, result[i].Name, expected[i].Name)
		}
		if result[i].Email != expected[i].Email {
			t.Errorf("MapAllUser() [%d] Email = %v, want %v", i, result[i].Email, expected[i].Email)
		}
		if result[i].CreatedAt != expected[i].CreatedAt {
			t.Errorf("MapAllUser() [%d] CreatedAt = %v, want %v", i, result[i].CreatedAt, expected[i].CreatedAt)
		}
		if result[i].UpdatedAt != expected[i].UpdatedAt {
			t.Errorf("MapAllUser() [%d] UpdatedAt = %v, want %v", i, result[i].UpdatedAt, expected[i].UpdatedAt)
		}
	}
}

func TestMapAllUserWithEmptySlice(t *testing.T) {
	// Test with empty slice
	testUsers := []domains.User{}
	result := dto.MapAllUser(testUsers)

	if len(result) != 0 {
		t.Errorf("MapAllUser() with empty slice length = %v, want %v", len(result), 0)
	}
}

func TestMapAllUserWithNilSlice(t *testing.T) {
	// Test with nil slice
	var testUsers []domains.User
	result := dto.MapAllUser(testUsers)

	if len(result) != 0 {
		t.Errorf("MapAllUser() with nil slice length = %v, want %v", len(result), 0)
	}
}
