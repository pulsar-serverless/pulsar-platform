package user

type GithubAuthReq struct {
	AccessToken string `json:"access_token"`
}

type GithubAuthResp struct {
	LoginName string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

func NewAuthReq(token string) *GithubAuthReq {
	return &GithubAuthReq{
		AccessToken: token,
	}
}
