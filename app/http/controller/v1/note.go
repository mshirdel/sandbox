package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
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
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NotesResponse struct {
	Notes []NoteResponse `json:"notes"`
}

type NoteController struct{}

func NewNoteController() *NoteController {
	return &NoteController{}
}

func (n *NoteController) GetNote(ctx echo.Context) error {
	req := new(GetNoteRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Mock response - in real app this would fetch from database
	note := NoteResponse{
		ID:      req.ID,
		Title:   "Sample Note",
		Content: "This is a sample note content",
	}

	return ctx.JSON(http.StatusOK, note)
}

func (n *NoteController) GetNotes(ctx echo.Context) error {
	// Mock response - in real app this would fetch from database
	notes := []NoteResponse{
		{ID: 1, Title: "First Note", Content: "Content of first note"},
		{ID: 2, Title: "Second Note", Content: "Content of second note"},
	}

	return ctx.JSON(http.StatusOK, NotesResponse{Notes: notes})
}

func (n *NoteController) CreateNote(ctx echo.Context) error {
	req := new(CreateNoteRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	logrus.Infof("Creating note with title: %s", req.Title)

	// Mock response - in real app this would save to database
	note := NoteResponse{
		ID:      123,
		Title:   req.Title,
		Content: req.Content,
	}

	return ctx.JSON(http.StatusCreated, note)
}

func (n *NoteController) UpdateNote(ctx echo.Context) error {
	req := new(UpdateNoteRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	logrus.Infof("Updating note ID: %d", req.ID)

	// Mock response - in real app this would update in database
	note := NoteResponse{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
	}

	return ctx.JSON(http.StatusOK, note)
}

func (n *NoteController) DeleteNote(ctx echo.Context) error {
	req := new(DeleteNoteRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	logrus.Infof("Deleting note ID: %d", req.ID)

	// Mock response - in real app this would delete from database
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Note deleted successfully"})
}
