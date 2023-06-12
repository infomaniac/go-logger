package gcplog

import (
	"testing"

	"github.com/infomaniac/go-logger"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*logger.ILogger)(nil), new(Logger))
}
