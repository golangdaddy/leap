type Users []*User

type UserRef struct {
	Mode     string
	ID       string
	Username string
}

func DemoUser() *User {
	return NewUser("demo", "john@doe.com", "john doe")
}

func NewUser(mode string, email, username string) *User {
	user := &User{
		Meta:     (Internals{}).NewInternals("users"),
		Mode:     mode,
		Email:    strings.ToLower(strings.TrimSpace(email)),
		Username: strings.ToLower(strings.TrimSpace(username)),
	}
	return user
}

type User struct {
	Meta     Internals
	Mode     string `json:"mode" firestore:"mode"`
	Email    string `json:"email" firestore:"email"`
	Username string `json:"username" firestore:"username"`
}

func (user *User) Ref() UserRef {
	return UserRef{
		Mode:     user.Mode,
		ID:       user.Meta.ID,
		Username: user.Username,
	}
}

func (users Users) Refs() []UserRef {
	refs := []UserRef{}
	for _, user := range users {
		refs = append(refs, user.Ref())
	}
	return refs
}

func (user *User) IsValid() bool {
	log.Println(user.Username)

	if len(user.Username) < 3 {
		return false
	}
	if len(user.Username) > 24 {
		return false
	}
	if strings.Contains(user.Username, " ") {
		return false
	}
	if !isAlphanumeric(strings.Replace(user.Username, "_", "", -1)) {
		return false
	}
	return true
}

func isAlphanumeric(word string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(word)
}

const (
	CONST_COL_SESSION = "sessions"
	CONST_COL_OTP     = "otp"
	CONST_COL_USER    = "users"
)

// GetOTP gets OTP record from firestore
func GetOTP(app *common.App, r *http.Request) (*OTP, error) {

	otp, err := cloudfunc.QueryParam(r, "otp")
	if err != nil {
		return nil, err
	}
	id := app.SeedDigest(otp)

	// fetch the OTP record
	doc, err := app.Firestore().Collection(CONST_COL_OTP).Doc(id).Get(app.Context())
	if err != nil {
		return nil, err
	}

	otpRecord := &OTP{}
	if err := doc.DataTo(&otpRecord); err != nil {
		return nil, err
	}

	// delete the OTP record
	if _, err := app.Firestore().Collection(CONST_COL_OTP).Doc(id).Delete(app.Context()); err != nil {
		return nil, err
	}

	return otpRecord, nil
}

// GetOTP gets OTP record from firestore
func (app *App) DebugGetOTP(r *http.Request) (*OTP, error) {

	otp, err := cloudfunc.QueryParam(r, "otp")
	if err != nil {
		return nil, err
	}
	id := app.SeedDigest(otp)

	// fetch the OTP record
	doc, err := app.Firestore().Collection(CONST_COL_OTP).Doc(id).Get(app.Context())
	if err != nil {
		return nil, err
	}

	otpRecord := &OTP{}
	if err := doc.DataTo(&otpRecord); err != nil {
		return nil, err
	}

	return otpRecord, nil
}

func (app *App) CreateSessionSecret(otp *OTP) (string, int64, error) {

	secret := app.Token256()
	hashedSecret := app.SeedDigest(secret)

	user, err := otp.GetUser(app.App)
	if err != nil {
		return "", 0, err
	}

	session := user.NewSession()

	// create the firestore session record
	if _, err := app.Firestore().Collection(CONST_COL_SESSION).Doc(hashedSecret).Set(app.Context(), session); err != nil {
		return "", 0, err
	}

	return secret, session.Expires, nil
}

func (app *App) GetSessionUser(r *http.Request) (*User, error) {

	apiKey := r.Header.Get("Authorization")
	if len(apiKey) == 0 {
		err := errors.New("missing apikey in Authorization header")
		return nil, err
	}
	id := app.SeedDigest(apiKey)

	// fetch the Session record
	doc, err := app.Firestore().Collection(CONST_COL_SESSION).Doc(id).Get(app.Context())
	if err != nil {
		return nil, err
	}
	session := &Session{}
	if err := doc.DataTo(&session); err != nil {
		return nil, err
	}

	// fetch the user record
	doc, err = app.Firestore().Collection(CONST_COL_USER).Doc(session.UserID).Get(app.Context())
	if err != nil {
		return nil, err
	}
	user := &User{}
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	return user, nil
}

// UserCollection abstracts the handling of subdata to within the user object
func (app *App) UserCollection(user *User, collectionID string) *firestore.CollectionRef {
	return app.UserRefCollection(user.Ref(), collectionID)
}

func (app *App) UserRefCollection(userRef UserRef, collectionID string) *firestore.CollectionRef {
	return app.Firestore().Collection("users").Doc(userRef.ID).Collection(collectionID)
}

// RegionCollection abstracts the handling of subdata to within the country/region
func (app *App) RegionCollection(user *User, collectionID string) *firestore.CollectionRef {
	return app.Firestore().Collection("countries").Doc(user.Meta.Context.Country).Collection("regions").Doc(user.Meta.Context.Region).Collection(collectionID)
}

func (app *App) GetUserByUsername(username string) (*User, error) {
	doc, err := app.Firestore().Collection("usernames").Doc(username).Get(app.Context())
	if err != nil {
		return nil, err
	}
	record := &Username{}
	if err := doc.DataTo(record); err != nil {
		return nil, err
	}
	return app.GetUserByID(record.User.ID)
}

func (app *App) GetUser(ref UserRef) (*User, error) {
	return app.GetUserByID(ref.ID)
}

func (app *App) GetUserByID(id string) (*User, error) {
	doc, err := app.Firestore().Collection("users").Doc(id).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := &User{}
	return user, doc.DataTo(user)
}

func (app *App) GetUserByEmail(email string) (*User, error) {

	iter := app.Firestore().Collection("users").Where("email", "==", email).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		user := &User{}
		return user, doc.DataTo(user)
	}

	return nil, fmt.Errorf("no user forund via email: %s", email)
}
