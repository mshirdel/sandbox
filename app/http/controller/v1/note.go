package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mshirdel/sandbox/models"
	"github.com/mshirdel/sandbox/repository"
	"github.com/sirupsen/logrus"
)

type GetNoteRequest struct {
	ID uint `param:"id" validate:"required"`
}

type CreateNoteRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdateNoteRequest struct {
	ID      uint   `param:"id" validate:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeleteNoteRequest struct {
	ID uint `param:"id" validate:"required"`
}

type NoteResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type NotesResponse struct {
	Notes []NoteResponse `json:"notes"`
}

type NoteController struct {
	noteRepository repository.NoteRepository
}

func NewNoteController(noteRepo repository.NoteRepository) *NoteController {
	return &NoteController{
		noteRepository: noteRepo,
	}
}

func (n *NoteController) GetNote(ctx echo.Context) error {
	req := new(GetNoteRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	note, err := n.noteRepository.GetByID(req.ID)
	if err != nil {
		logrus.Errorf("Failed to get note with ID %d: %v", req.ID, err)
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Note not found"})
	}

	response := NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: note.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return ctx.JSON(http.StatusOK, response)
}

func (n *NoteController) GetNotes(ctx echo.Context) error {
	notes, err := n.noteRepository.GetAll()
	if err != nil {
		logrus.Errorf("Failed to get notes: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve notes"})
	}

	var responses []NoteResponse
	for _, note := range notes {
		responses = append(responses, NoteResponse{
			ID:        note.ID,
			Title:     note.Title,
			Content:   note.Content,
			CreatedAt: note.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt: note.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return ctx.JSON(http.StatusOK, NotesResponse{Notes: responses})
}

func (n *NoteController) CreateNote(ctx echo.Context) error {
	req := new(CreateNoteRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	note := &models.Note{
		Title:   req.Title,
		Content: req.Content,
	}

	err := n.noteRepository.Create(note)
	if err != nil {
		logrus.Errorf("Failed to create note: %v", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create note"})
	}

	logrus.Infof("Created note with ID: %d", note.ID)

	response := NoteResponse{
		ID:        note.ID,
		Title:     note.Title,
		Content:   note.Content,
		CreatedAt: note.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: note.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (n *NoteController) UpdateNote(ctx echo.Context) error {
	req := new(UpdateNoteRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// First get the existing note
	existingNote, err := n.noteRepository.GetByID(req.ID)
	if err != nil {
		logrus.Errorf("Failed to get note with ID %d: %v", req.ID, err)
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Note not found"})
	}

	// Update fields if provided
	if req.Title != "" {
		existingNote.Title = req.Title
	}
	if req.Content != "" {
		existingNote.Content = req.Content
	}

	err = n.noteRepository.Update(existingNote)
	if err != nil {
		logrus.Errorf("Failed to update note with ID %d: %v", req.ID, err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update note"})
	}

	logrus.Infof("Updated note ID: %d", req.ID)

	response := NoteResponse{
		ID:        existingNote.ID,
		Title:     existingNote.Title,
		Content:   existingNote.Content,
		CreatedAt: existingNote.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: existingNote.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	return ctx.JSON(http.StatusOK, response)
}

func (n *NoteController) DeleteNote(ctx echo.Context) error {
	req := new(DeleteNoteRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := n.noteRepository.Delete(req.ID)
	if err != nil {
		logrus.Errorf("Failed to delete note with ID %d: %v", req.ID, err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete note"})
	}

	logrus.Infof("Deleted note ID: %d", req.ID)

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Note deleted successfully"})
}
