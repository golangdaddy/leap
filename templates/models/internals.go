type Generic struct {
	Meta Internals
}

func Internal(id string) *Internals {
	return &Internals{ID: id}
}

// NewInternals returns a boilerplate internal object
func (n Internals) NewInternals(class string) Internals {

	timestamp := time.Now().UTC().Unix()

	x := Internals{
		ID:       n.ID + "." + class + "-" + uuid.NewString()[:13],
		Class:    class,
		Created:  timestamp,
		Modified: timestamp,
	}
	if len(n.ID) > 0 {
		x.Context.Parent = n.ID
		x.Context.Parents = append(n.Context.Parents, n.ID)
	}
	return x
}

type Internals struct {
	ID         string
	Class      string
	ClassName  string
	Name       string `json:",omitempty"`
	Asset      string `json:",omitempty"`
	Wallet     string `json:",omitempty"`
	Context    Context
	Moderation Moderation
	Media      Media
	Updated    bool
	Created    int64
	Deleted    int64 `json:",omitempty"`
	Modified   int64
	Stats      map[string]float64 `json:",omitempty"`
}

func RegExp(exp, matchString string) bool {
	return regexp.MustCompile(exp).MatchString(matchString)
}

func (i *Internals) Assetlayer() *assetlayer.Client {
	return assetlayer.NewClient(
		os.Getenv("APPID"),
		os.Getenv("ASSETLAYERSECRET"),
		os.Getenv("DIDTOKEN"),
	)
}

func (i *Internals) AssetlayerWalletID() string {
	x := sha256.Sum256([]byte(i.ID))
	return hex.EncodeToString(x[:])[:32]
}

func (i *Internals) AssetlayerCollectionID() string {
	return os.Getenv("MODEL_" + strings.ToUpper(i.Class))
}

func (i *Internals) URI() (string, error) {
	if len(i.Media.URIs) == 0 {
		return "", errors.New("this object has no assigned URI")
	}
	return i.Media.URIs[len(i.Media.URIs)-1], nil
}

func (i *Internals) NewURI() string {
	i.Media.URIs = append(i.Media.URIs, uuid.NewString())
	i.Modify()
	return i.Media.URIs[len(i.Media.URIs)-1]
}

func (i *Internals) DocPath() string {
	println("docpath:", i.ID)
	p := strings.Split(string(i.ID[1:]), ".")
	parts := make([][]string, len(p))
	k := ""
	for x, s := range p {
		k += "." + s
		parts[x] = strings.Split(k, ".")

	}
	outs := []string{}
	for _, p := range parts {
		class := strings.Split(p[len(p)-1], "-")[0]
		outs = append(outs, class+"/"+strings.Join(p, "."))
	}
	return strings.Join(outs, "/")
}

func (i *Internals) SaveToFirestore(app *common.App, src interface{}) error {
	i.Modify()
	_, err := i.Firestore(app).Set(app.Context(), src)
	return err
}

func (i *Internals) Firestore(app *common.App) *firestore.DocumentRef {
	path := i.DocPath()
	return app.Firestore().Doc(path)
}

func (i *Internals) FirestoreDoc(app *common.App, ii Internals) *firestore.DocumentRef {
	return i.Firestore(app).Collection(ii.Class).Doc(ii.ID)
}

func FirestoreCount(app *common.App, collection string) int {
	var count int
	iter := app.Firestore().Collection(collection).Documents(app.Context())
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		count++
	}
	return count
}

func (i *Internals) FirestoreCount(app *common.App, collection string) int {
	var count int
	iter := i.Firestore(app).Collection(collection).Documents(app.Context())
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		count++
	}
	return count
}

func ClassFromID(id string) (string, error) {
	p := strings.Split(id, ".")
	last := p[len(p)-1]
	class := strings.Split(last, "-")
	if len(class) != 2 {
		return "", fmt.Errorf("CANT GET CLASS FROM ID: %s", id)
	}
	return class[0], nil
}

func (i *Internals) ParentID() (string, error) {
	if len(i.Context.Parents) == 0 {
		return "", fmt.Errorf("%s has no parents", i.Class)
	}
	return i.Context.Parents[len(i.Context.Parents)-1], nil
}

func (i *Internals) GetParent(app *common.App, dst interface{}) error {
	parentID, err := i.ParentID()
	if err != nil {
		return err
	}
	parent := Internal(parentID)
	doc, err := parent.Firestore(app).Get(app.Context())
	if err != nil {
		return err
	}
	return doc.DataTo(dst)
}

func (i *Internals) GetParentMeta(app *common.App, dst interface{}) error {
	parentID, err := i.ParentID()
	if err != nil {
		return err
	}
	parent := Internal(parentID)
	doc, err := parent.Firestore(app).Get(app.Context())
	if err != nil {
		return err
	}
	return doc.DataTo(dst)
}

// Modify updates the timestamp
func (i *Internals) Modify() {
	i.Modified = time.Now().UTC().Unix()
}

// Update sets the metadata to indicate it has updated for a user
func (i *Internals) Update() {
	i.Updated = true
	i.Modify()
}

type Media struct {
	Color  string                 `json:",omitempty"`
	URIs   []string               `json:",omitempty"`
	Image  bool                   `json:",omitempty"`
	EXIF   map[string]interface{} `json:",omitempty"`
	Format string                 `json:",omitempty"`
}

type Context struct {
	Children []string `json:",omitempty"`
	Parent   string   `json:",omitempty"`
	Parents  []string `json:",omitempty"`
	Country  string   `json:",omitempty"`
	Region   string   `json:",omitempty"`
	Order    int
	Status   string `json:",omitempty"`
}

type Moderation struct {
	// used for an object id too inherit admins for
	Object string `json:",omitempty"`
	// list of admin usernames
	Admins       []string `json:",omitempty"`
	Blocked      bool     `json:",omitempty"`
	BlockedTime  int64    `json:",omitempty"`
	BlockedBy    string   `json:",omitempty"`
	Approved     bool     `json:",omitempty"`
	ApprovedTime int64    `json:",omitempty"`
	ApprovedBy   string   `json:",omitempty"`
}
