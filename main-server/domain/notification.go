package domain

type Notification struct {
	FromUser UserId
	ToUser   UserId
	Text     string
}
