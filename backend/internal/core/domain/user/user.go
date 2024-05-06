package user

type User struct {
	UserId       string `json:"userId"`
	ProjectCount int    `json:"projectCount"`
}

func New() User {
	return User{}
}
