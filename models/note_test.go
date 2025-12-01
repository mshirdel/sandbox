package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_Note_TableName(t *testing.T) {
	note := Note{}
	assert.Equal(t, "notes", note.TableName())
}

func Test_Note_Fields(t *testing.T) {
	now := time.Now()
	note := Note{
		ID:        1,
		Title:     "Test Note",
		Content:   "Test Content",
		CreatedAt: now,
		UpdatedAt: now,
	}

	assert.Equal(t, uint(1), note.ID)
	assert.Equal(t, "Test Note", note.Title)
	assert.Equal(t, "Test Content", note.Content)
	assert.Equal(t, now, note.CreatedAt)
	assert.Equal(t, now, note.UpdatedAt)
	assert.IsType(t, gorm.DeletedAt{}, note.DeletedAt)
}
