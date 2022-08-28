package models

import (
	"testing"

	"snippetbox.umaralfaruq/internal/assert"
)

func TestSnippetModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	var (
		snippetID   = 1
		wantID      = 1
		wantTitle   = "cat1"
		wantContent = "pompomtut1"
	)

	t.Run("Valid ID", func(t *testing.T) {
		db := newTestDB(t)

		m := SnippetModel{db}

		snippet, err := m.Get(snippetID)

		assert.Equal(t, snippet.ID, wantID)
		assert.Equal(t, snippet.Title, wantTitle)
		assert.Equal(t, snippet.Content, wantContent)
		assert.NilError(t, err)
	})
}
