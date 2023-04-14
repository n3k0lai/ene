package Users

type User struct {
	Username    string
	Displayname string
	AvatarUrl   string
}

func NewUser() *User {
	return &User{}
}
