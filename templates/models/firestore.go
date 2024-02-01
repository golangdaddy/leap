func (app *App) IsAdmin(parent *Internals, user *User) bool {
	if len(parent.Moderation.Object) > 0 {
		var err error
		parent, err = app.GetMetadata(parent.Moderation.Object)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	for _, userID := range parent.Moderation.Admins {
		if user.Meta.ID == userID {
			return true
		}
	}
	return false
}

func (app *App) GetMetadata(id string) (*Internals, error) {

	dst := &Generic{}

	i := Internal(id)
	path := i.DocPath()

	println("GET DOCUMENT", path)

	doc, err := app.Firestore().Doc(path).Get(context.Background())
	if err != nil {
		return nil, err
	}
	return &dst.Meta, doc.DataTo(dst)
}

func (app *App) GetDocument(id string, dst interface{}) error {

	i := Internal(id)
	path := i.DocPath()

	println("GET DOCUMENT", path)

	doc, err := app.Firestore().Doc(path).Get(context.Background())
	if err != nil {
		return err
	}
	return doc.DataTo(dst)
}
