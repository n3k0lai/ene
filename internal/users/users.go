package Users

type User struct {
	Username    string
	Displayname string
	AvatarUrl   string
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}
