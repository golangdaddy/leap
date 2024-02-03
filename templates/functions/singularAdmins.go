package main

import (
	"log"
	"strings"
)

func (app *App) add{{titlecase .Object.Name}}Admin(object *{{uppercase .Object.Name}}, admin string) error {

	filter := map[string]bool{}
	for _, username := range strings.Split(admin, ",") {
		newAdmin, err := app.GetUserByUsername(username)
		if err != nil {
			log.Println("could not get username:", username)
			return err
		}
		filter[newAdmin.Meta.ID] = true
	}
	for _, admin := range object.Meta.Moderation.Admins {
		if len(admin) == 0 {
			continue
		}
		filter[admin] = true
	}
	object.Meta.Moderation.Admins = make([]string, len(filter))
	var x int
	for k, _ := range filter {
		object.Meta.Moderation.Admins[x] = k
		x++
	}

	object.Meta.Modify()

	log.Println("ADMINS", strings.Join(object.Meta.Moderation.Admins, " "))

	return object.Meta.SaveToFirestore(app.App, object)
}

func (app *App) remove{{titlecase .Object.Name}}Admin(object *{{uppercase .Object.Name}}, admin string) error {

	filter := map[string]bool{}
	for _, a := range object.Meta.Moderation.Admins {
		if a == admin {
			continue
		}
		if len(a) == 0 {
			continue
		}
		filter[a] = true
	}
	object.Meta.Moderation.Admins = make([]string, len(filter))
	var x int
	for k, _ := range filter {
		object.Meta.Moderation.Admins[x] = k
		x++
	}

	object.Meta.Modify()

	log.Println("ADMINS", strings.Join(object.Meta.Moderation.Admins, " "))

	return object.Meta.SaveToFirestore(app.App, object)
}