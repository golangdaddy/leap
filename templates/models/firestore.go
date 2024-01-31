
func GetMetadata(app *common.App, id string) (*Internals, error) {

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

func GetDocument(app *common.App, id string, dst interface{}) error {

	i := Internal(id)
	path := i.DocPath()

	println("GET DOCUMENT", path)

	doc, err := app.Firestore().Doc(path).Get(context.Background())
	if err != nil {
		return err
	}
	return doc.DataTo(dst)
}
