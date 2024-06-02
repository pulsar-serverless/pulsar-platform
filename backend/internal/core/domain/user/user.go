package user

type User struct {
	UserId       string `json:"userId"`
	ProjectCount int    `json:"projectCount"`
	Status       string `json:"status"`
}

func New() User {
	return User{}
}
