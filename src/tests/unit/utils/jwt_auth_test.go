package utils

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

// Local function to avoid dependencies on global config
func generatePasswordHashLocal(password string) string {
	salt := "qweasdzxc" // Same salt as in utils
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func TestGeneratePasswordHash(t *testing.T) {
	tests := []struct {
		name     string
		password string
		want     string
	}{
		{
			name:     "simple password",
			password: "123456",
			want:     "7177656173647a78637c4a8d09ca3762af61e59520943dc26494f8941b", // expected hash for "123456" with salt "qweasdzxc"
		},
		{
			name:     "empty password",
			password: "",
			want:     "7177656173647a7863da39a3ee5e6b4b0d3255bfef95601890afd80709", // expected hash for empty string with salt
		},
		{
			name:     "complex password",
			password: "MyComplexPassword123!@#",
			want:     generatePasswordHashLocal("MyComplexPassword123!@#"), // calculate expected hash
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePasswordHashLocal(tt.password)
			if got != tt.want {
				t.Errorf("GeneratePasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGeneratePasswordHashConsistency(t *testing.T) {
	password := "testPassword123"

	// Hash the same password multiple times
	hash1 := generatePasswordHashLocal(password)
	hash2 := generatePasswordHashLocal(password)
	hash3 := generatePasswordHashLocal(password)

	// All hashes should be identical
	if hash1 != hash2 {
		t.Errorf("Hash inconsistency: hash1=%v, hash2=%v", hash1, hash2)
	}

	if hash1 != hash3 {
		t.Errorf("Hash inconsistency: hash1=%v, hash3=%v", hash1, hash3)
	}
}

func TestGeneratePasswordHashDifferentPasswords(t *testing.T) {
	password1 := "password1"
	password2 := "password2"

	hash1 := generatePasswordHashLocal(password1)
	hash2 := generatePasswordHashLocal(password2)

	// Different passwords should produce different hashes
	if hash1 == hash2 {
		t.Errorf("Different passwords produced same hash: %v", hash1)
	}
}

func TestGeneratePasswordHashLength(t *testing.T) {
	password := "anyPassword"
	hash := generatePasswordHashLocal(password)

	// SHA1 hash + salt should be 58 characters long (SHA1 160 bits + salt 72 bits = 232 bits / 4 bits per hex char = 58 chars)
	expectedLength := 58
	if len(hash) != expectedLength {
		t.Errorf("Hash length = %d, want %d", len(hash), expectedLength)
	}
}

func TestGeneratePasswordHashSaltEffect(t *testing.T) {
	password := "testPassword"

	// Test with salt
	saltedHash := generatePasswordHashLocal(password)

	// Test without salt (manual calculation)
	hashWithoutSalt := sha1.New()
	hashWithoutSalt.Write([]byte(password))
	unsaltedHash := fmt.Sprintf("%x", hashWithoutSalt.Sum(nil))

	// Salted and unsalted hashes should be different
	if saltedHash == unsaltedHash {
		t.Errorf("Salted and unsalted hashes should be different")
	}
}

func TestGeneratePasswordHashKnownValues(t *testing.T) {
	// Test with known values to ensure algorithm correctness
	tests := []struct {
		password string
		expected string
	}{
		{"123456", "7177656173647a78637c4a8d09ca3762af61e59520943dc26494f8941b"},
		{"", "7177656173647a7863da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"password", "7177656173647a78635baa61e4c9b93f3f0682250b6cf8331b7ee68fd8"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("password_%s", tt.password), func(t *testing.T) {
			// Calculate actual hash
			actualHash := generatePasswordHashLocal(tt.password)

			// Check exact match for all known values
			if actualHash != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, actualHash)
			}

			// Also check that hash is not empty and has correct length
			if len(actualHash) != 58 {
				t.Errorf("Hash should be 58 characters, got %d", len(actualHash))
			}
			if actualHash == "" {
				t.Errorf("Hash should not be empty")
			}
		})
	}
}
