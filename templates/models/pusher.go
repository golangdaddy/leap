func (app *App) SendMessageToUser(user *User, msgType string, data interface{}) {
	log.Println("SENDING MESSAGE TO PUSHER USER:", user.Username)
	err := app.Pusher().Trigger(user.Meta.ID, msgType, data)
	if err != nil {
		log.Println(err.Error())
	}
}
