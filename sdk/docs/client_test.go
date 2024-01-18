package docs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	assert := assert.New(t)

	client := NewClient()

	assert.NotNil(client)
}
