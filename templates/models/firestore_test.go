package main

import (
	"testing"

	"github.com/richardboase/npgpublic/models"
)

func TestPath(t *testing.T) {

	{
		internals := models.Internals{}
		internals.ID = ".projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1.collections-b9416363-eaf6-4e41-9ece-9ffc0026b289"

		result := (&internals).DocPath()
		expected := "projects/.projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1/collections/.projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1.collections-b9416363-eaf6-4e41-9ece-9ffc0026b289"
		if result != expected {
			t.Fatal()
		}
	}

	{
		internals := models.Internals{}
		internals.ID = ".projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1.collections-b9416363-eaf6-4e41-9ece-9ffc0026b289.cunts-b941636-eaf6-4e41-9ece-9ffc0026b289"

		result := (&internals).DocPath()
		expected := "projects/.projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1/collections/.projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1.collections-b9416363-eaf6-4e41-9ece-9ffc0026b289/cunts/.projects-e544b82a-f5ca-4dee-9722-09f80e50d4e1.collections-b9416363-eaf6-4e41-9ece-9ffc0026b289.cunts-b941636-eaf6-4e41-9ece-9ffc0026b289"
		if result != expected {
			t.Fatal()
		}
	}
}
