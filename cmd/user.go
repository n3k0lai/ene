package core

type User struct {
	Username string
	Displayname string
	AvatarUrl string
	AdapterType AdapterType
}

func NewUser() *User {
	return &User{}
}
