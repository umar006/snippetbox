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
