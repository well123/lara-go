package schema

type LoginParam struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginUserParam struct {
	ApiId     string `json:"apiId" validator:"required"`
	ApiSecret string `json:"apiSecret" validator:"required,max=32"`
}

type LoginTypeInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
	AccessType  string `json:"access_type"`
}
