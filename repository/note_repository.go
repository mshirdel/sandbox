package repository

import (
	"github.com/mshirdel/sandbox/models"
	"gorm.io/gorm"
)

// NoteRepository defines the interface for note data operations
type NoteRepository interface {
	Create(note *models.Note) error
	GetByID(id uint) (*models.Note, error)
	GetAll() ([]*models.Note, error)
	Update(note *models.Note) error
	Delete(id uint) error
}

// noteRepository implements NoteRepository interface
type noteRepository struct {
	db *gorm.DB
}

// NewNoteRepository creates a new note repository instance
func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{
		db: db,
	}
}

// Create creates a new note
func (r *noteRepository) Create(note *models.Note) error {
	return r.db.Create(note).Error
}

// GetByID retrieves a note by ID
func (r *noteRepository) GetByID(id uint) (*models.Note, error) {
	var note models.Note
	err := r.db.First(&note, id).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

// GetAll retrieves all notes
func (r *noteRepository) GetAll() ([]*models.Note, error) {
	var notes []*models.Note
	err := r.db.Find(&notes).Error
	return notes, err
}

// Update updates an existing note
func (r *noteRepository) Update(note *models.Note) error {
	return r.db.Save(note).Error
}

// Delete soft deletes a note by ID
func (r *noteRepository) Delete(id uint) error {
	return r.db.Delete(&models.Note{}, id).Error
}
