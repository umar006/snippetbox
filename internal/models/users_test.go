package models

import (
	"testing"

	"snippetbox.umaralfaruq/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "Valid ID",
			userID: 1,
			want:   true,
		},
		{
			name:   "Zero ID",
			userID: 0,
			want:   false,
		},
		{
			name:   "Non-existent ID",
			userID: 2,
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			exists, err := m.Exists(test.userID)

			assert.Equal(t, exists, test.want)
			assert.NilError(t, err)
		})
	}
}

func TestUserModelInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name         string
		userID       int
		userName     string
		userEmail    string
		userPassword string
		want         bool
	}{
		{
			name:         "Valid insert",
			userID:       2,
			userName:     "umar",
			userEmail:    "umar@test.com",
			userPassword: "test",
			want:         true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			err := m.Insert(test.userName, test.userEmail, test.userPassword)

			assert.NilError(t, err)

			exists, err := m.Exists(test.userID)

			assert.Equal(t, exists, test.want)
			assert.NilError(t, err)
		})
	}
}

func TestUserModelInsertError(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name         string
		userID       int
		userName     string
		userEmail    string
		userPassword string
		wantExists   bool
		wantError    error
	}{
		{
			name:         "Duplicate email",
			userID:       2,
			userName:     "Alice",
			userEmail:    "alice@example.com",
			userPassword: "test",
			wantExists:   false,
			wantError:    ErrDuplicateEmail,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			err := m.Insert(test.userName, test.userEmail, test.userPassword)

			assert.EqualError(t, err, test.wantError)

			exists, err := m.Exists(test.userID)

			assert.Equal(t, exists, test.wantExists)
			assert.NilError(t, err)
		})
	}
}

func TestUserModelAuthenticate(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name         string
		userEmail    string
		userPassword string
		wantID       int
	}{
		{
			name:         "User authenticated",
			userEmail:    "alice@example.com",
			userPassword: "alice",
			wantID:       1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			id, err := m.Authenticate(test.userEmail, test.userPassword)

			assert.NilError(t, err)
			assert.Equal(t, id, test.wantID)
		})
	}
}

func TestUserModelAuthenticateError(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name         string
		userEmail    string
		userPassword string
		wantID       int
		wantError    error
	}{
		{
			name:         "Incorrect email",
			userEmail:    "umar@example.com",
			userPassword: "alice",
			wantID:       0,
			wantError:    ErrInvalidCredentials,
		},
		{
			name:         "Incorrect password",
			userEmail:    "alice@example.com",
			userPassword: "umaru",
			wantID:       0,
			wantError:    ErrInvalidCredentials,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			id, err := m.Authenticate(test.userEmail, test.userPassword)

			assert.EqualError(t, err, test.wantError)
			assert.Equal(t, id, test.wantID)
		})
	}
}
