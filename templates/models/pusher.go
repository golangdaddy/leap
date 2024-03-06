func (app *App) SendMessageToUser(user *User, msgType string, data interface{}) {
	{{if .Options.Pusher}}
	log.Printf("SENDING %s MESSAGE TO PUSHER USER %s (%s)", msgType, user.Username, user.Meta.ID)
	err := app.Pusher().Trigger(user.Meta.ID, msgType, data)
	if err != nil {
		log.Println(err.Error())
	}
	{{end}}
}
