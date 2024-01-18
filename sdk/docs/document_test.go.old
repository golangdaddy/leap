package docs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocument(t *testing.T) {
	assert := assert.New(t)

	where := NewPlace("FR")

	what := map[string]interface{}{
		"hello": "world",
	}

	client := NewClient()
	doc := client.NewDocument(where, "demoobject", what)

	err := doc.Save()
	assert.Nil(err)
}
