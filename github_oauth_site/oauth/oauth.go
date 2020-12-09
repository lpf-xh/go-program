package oauth

// OAuth 授权接口
type OAuth interface {
	SetCode(code string)
	GetToken() (err error)
	GetUser() (name string, err error)
}
