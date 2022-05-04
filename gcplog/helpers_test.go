package gcplog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentProject(t *testing.T) {
	_, err := CurrentProject()
	assert.NoError(t, err)
}
