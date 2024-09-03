package project

import (
	"github.com/golangdaddy/leap/models"
)

func buildStructure(config models.Config) *models.Stack {

	tree := &models.Stack{
		WebsiteName: "Hello World",
		Config:      config,
		Options:     models.StackOptions{},
	}

	return tree
}
