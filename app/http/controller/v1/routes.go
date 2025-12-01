package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/mshirdel/sandbox/app"
)

type Router struct {
	app            *app.Application
	userController *UserController
	noteController *NoteController
}

func New(a *app.Application) *Router {
	return &Router{
		app:            a,
		userController: NewUserController(),
		noteController: NewNoteController(a.NoteRepository),
	}
}

func (r *Router) Routes(g *echo.Group) {
	user := g.Group("/users")
	{
		user.GET("/:id", r.userController.GetUser)
		user.POST("", r.userController.SaveMessage)
	}

	note := g.Group("/notes")
	{
		note.GET("", r.noteController.GetNotes)
		note.GET("/:id", r.noteController.GetNote)
		note.POST("", r.noteController.CreateNote)
		note.PUT("/:id", r.noteController.UpdateNote)
		note.DELETE("/:id", r.noteController.DeleteNote)
	}
}
