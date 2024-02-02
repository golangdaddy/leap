type Mail struct {
	Meta       Internals
	Sender     UserRef   `json:"sender" firestore:"sender"`
	Recipients []UserRef `json:"recipients" firestore:"recipients"`
	Subject    string    `json:"subject" firestore:"subject"`
	Body       string    `json:"body" firestore:"body"`
}

func (user *User) NewMail(subject, body string, recipients ...*User) *Mail {
	return &Mail{
		Meta:       user.Meta.NewInternals("mail"),
		Sender:     user.Ref(),
		Recipients: Users(recipients).Refs(),
		Subject:    subject,
		Body:       body,
	}
}

type MailReply struct {
	Meta       Internals
	ID         string    `json:"id" firestore:"id"`
	Sender     UserRef   `json:"sender" firestore:"sender"`
	Recipients []UserRef `json:"recipients" firestore:"recipients"`
	Subject    string    `json:"subject" firestore:"subject"`
	Body       string    `json:"body" firestore:"body"`
}

func (user *User) NewMailReply(op *Mail, body string, additionalRecipients ...UserRef) *MailReply {
	mail := &MailReply{
		Meta:       user.Meta.NewInternals("mailreply"),
		ID:         uuid.NewString(),
		Sender:     user.Ref(),
		Recipients: append(op.Recipients, op.Sender),
		Subject:    op.Subject,
		Body:       body,
	}
	mail.Meta.Context.Parent = op.Meta.ID
	// ensure no duplicate recipients
	filter := map[string]UserRef{}
	// merge existing recipients with additional ones
	for _, r := range append(mail.Recipients, additionalRecipients...) {
		filter[r.ID] = r
	}
	mail.Recipients = make([]UserRef, len(filter))
	var n int
	for _, recipient := range filter {
		mail.Recipients[n] = recipient
		n++
	}
	return mail
}
