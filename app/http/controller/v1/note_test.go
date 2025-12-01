package v1_test

import (
	"testing"

	v1 "github.com/mshirdel/sandbox/app/http/controller/v1"
	"github.com/stretchr/testify/assert"
)

func Test_NewNoteController(t *testing.T) {
	controller := v1.NewNoteController()
	assert.NotNil(t, controller)
}

func Test_GetNotes_Success(t *testing.T) {
	controller := v1.NewNoteController()
	assert.NotNil(t, controller)
	// Additional test logic would go here with HTTP testing
}

func Test_CreateNote_Success(t *testing.T) {
	controller := v1.NewNoteController()
	assert.NotNil(t, controller)
	// Additional test logic would go here with HTTP testing
}
