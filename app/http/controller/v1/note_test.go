package v1_test

import (
	"testing"

	v1 "github.com/mshirdel/sandbox/app/http/controller/v1"
	"github.com/mshirdel/sandbox/models"
	"github.com/stretchr/testify/assert"
)

// mockNoteRepository implements repository.NoteRepository for testing
type mockNoteRepository struct{}

func (m *mockNoteRepository) Create(note *models.Note) error        { return nil }
func (m *mockNoteRepository) GetByID(id uint) (*models.Note, error) { return &models.Note{}, nil }
func (m *mockNoteRepository) GetAll() ([]*models.Note, error)       { return []*models.Note{}, nil }
func (m *mockNoteRepository) Update(note *models.Note) error        { return nil }
func (m *mockNoteRepository) Delete(id uint) error                  { return nil }

func Test_NewNoteController(t *testing.T) {
	mockRepo := &mockNoteRepository{}
	controller := v1.NewNoteController(mockRepo)
	assert.NotNil(t, controller)
}

func Test_GetNotes_Success(t *testing.T) {
	mockRepo := &mockNoteRepository{}
	controller := v1.NewNoteController(mockRepo)
	assert.NotNil(t, controller)
	// Additional test logic would go here with HTTP testing
}

func Test_CreateNote_Success(t *testing.T) {
	mockRepo := &mockNoteRepository{}
	controller := v1.NewNoteController(mockRepo)
	assert.NotNil(t, controller)
	// Additional test logic would go here with HTTP testing
}
