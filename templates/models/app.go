//

type App struct {
	*common.App
	connections map[string]*websocket.Conn
	sync.RWMutex
}

func NewApp() *App {
	app := &App{
		App:         common.NewApp(),
		connections: map[string]*websocket.Conn{},
	}
	app.UseGCP(CONST_PROJECT_ID)
	app.UseGCPFirestore(CONST_FIRESTORE_DB)
	return app
}

func PrettyPrint(data interface{}) {
	bytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("Error pretty printing data: %s", err)
	}
	fmt.Println(string(bytes))
}
