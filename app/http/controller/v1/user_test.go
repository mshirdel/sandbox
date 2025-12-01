package v1_test

import (
	"testing"

	v1 "github.com/mshirdel/sandbox/app/http/controller/v1"
	"github.com/stretchr/testify/assert"
)

func Test_GetUser_Success(t *testing.T) {
	err := v1.Add(42)
	assert.NoError(t, err)
}
