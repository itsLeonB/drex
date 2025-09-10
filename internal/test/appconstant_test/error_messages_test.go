package appconstant_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/appconstant"
	"github.com/stretchr/testify/assert"
)

func TestErrorMessages_Constants(t *testing.T) {
	assert.Equal(t, "error retrieving data", appconstant.ErrDataSelect)
	assert.Equal(t, "transfer method with ID: %s is not found", appconstant.ErrTransferMethodNotFound)
	assert.Equal(t, "transfer method with ID: %s is deleted", appconstant.ErrTransferMethodDeleted)
}
