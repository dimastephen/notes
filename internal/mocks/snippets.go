package mocks

import (
	"github.com/dimastephen/snippetbox/internal/models"
	"time"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "Something",
	Content: "BUBABUBAIII",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 5, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
