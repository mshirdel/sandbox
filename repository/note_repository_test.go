package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/mshirdel/sandbox/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB creates an in-memory SQLite database for testing
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Auto migrate the schema
	err = db.AutoMigrate(&models.Note{})
	assert.NoError(t, err)

	return db
}

func Test_NewNoteRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := NewNoteRepository(db)
	assert.NotNil(t, repo)
}

func Test_NoteRepository_Create(t *testing.T) {
	db := setupTestDB(t)
	repo := NewNoteRepository(db)

	note := &models.Note{
		Title:   "Test Note",
		Content: "Test Content",
	}

	err := repo.Create(note)
	assert.NoError(t, err)
	assert.NotZero(t, note.ID)
	assert.NotZero(t, note.CreatedAt)
	assert.NotZero(t, note.UpdatedAt)
}

func Test_NoteRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewNoteRepository(db)

	// Create a note first
	note := &models.Note{
		Title:   "Test Note",
		Content: "Test Content",
	}
	err := repo.Create(note)
	assert.NoError(t, err)

	// Retrieve the note
	retrieved, err := repo.GetByID(note.ID)
	assert.NoError(t, err)
	assert.Equal(t, note.ID, retrieved.ID)
	assert.Equal(t, note.Title, retrieved.Title)
	assert.Equal(t, note.Content, retrieved.Content)
}

func Test_NoteRepository_GetByID_NotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewNoteRepository(db)

	_, err := repo.GetByID(999)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func Test_NoteRepository_GetAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewNoteRepository(db)

	// Create multiple notes
	notes := []*models.Note{
		{Title: "Note 1", Content: "Content 1"},
		{Title: "Note 2", Content: "Content 2"},
		{Title: "Note 3", Content: "Content 3"},
	}

	for _, note := range notes {
		err := repo.Create(note)
		assert.NoError(t, err)
	}

	// Retrieve all notes
	allNotes, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, allNotes, 3)
}

func Test_NoteRepository_Update(t *testing.T) {
	db := setupTestDB(t)
	repo := NewNoteRepository(db)

	// Create a note
	note := &models.Note{
		Title:   "Original Title",
		Content: "Original Content",
	}
	err := repo.Create(note)
	assert.NoError(t, err)

	originalUpdatedAt := note.UpdatedAt

	// Wait a bit to ensure updated_at changes
	time.Sleep(time.Millisecond)

	// Update the note
	note.Title = "Updated Title"
	note.Content = "Updated Content"
	err = repo.Update(note)
	assert.NoError(t, err)

	// Verify the update
	assert.Equal(t, "Updated Title", note.Title)
	assert.Equal(t, "Updated Content", note.Content)
	assert.True(t, note.UpdatedAt.After(originalUpdatedAt))
}

func Test_NoteRepository_Delete(t *testing.T) {
	db := setupTestDB(t)
	repo := NewNoteRepository(db)

	// Create a note
	note := &models.Note{
		Title:   "Test Note",
		Content: "Test Content",
	}
	err := repo.Create(note)
	assert.NoError(t, err)

	// Delete the note
	err = repo.Delete(note.ID)
	assert.NoError(t, err)

	// Verify it's soft deleted (should not be found in normal queries)
	_, err = repo.GetByID(note.ID)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
