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

func TestSnippetModelGetError(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name      string
		snippetID int
		wantError error
	}{
		{
			name:      "Zero ID",
			snippetID: 0,
			wantError: ErrNoRecord,
		},
		{
			name:      "Negative ID",
			snippetID: -1,
			wantError: ErrNoRecord,
		},
		{
			name:      "Non-existent ID",
			snippetID: 5,
			wantError: ErrNoRecord,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := newTestDB(t)

			m := SnippetModel{db}

			snippet, err := m.Get(test.snippetID)

			assert.EqualError(t, err, test.wantError)
			assert.Equal(t, snippet, nil)
		})
	}
}

func TestSnippetModelInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	var (
		title   = "cat4"
		content = "poempoemtut4"
		expires = 7
		wantID  = 4
	)

	t.Run("Valid snippet", func(t *testing.T) {
		db := newTestDB(t)

		m := SnippetModel{db}

		id, err := m.Insert(title, content, expires)

		assert.NilError(t, err)
		assert.Equal(t, id, wantID)
	})
}

func TestSnippetModelLatest(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	wantSnippet := []Snippet{
		{
			ID:      1,
			Title:   "cat1",
			Content: "poempoemtut1",
		},
		{
			ID:      2,
			Title:   "cat2",
			Content: "poempoemtut2",
		},
		{
			ID:      3,
			Title:   "cat3",
			Content: "poempoemtut3",
		},
	}

	t.Run("Get 3 snippets", func(t *testing.T) {
		db := newTestDB(t)

		m := SnippetModel{db}

		snippets, err := m.Latest()

		assert.NilError(t, err)
		assert.Equal(t, len(snippets), len(wantSnippet))
	})
}
