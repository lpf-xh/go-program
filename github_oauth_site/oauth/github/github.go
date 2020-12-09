package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// clientid 和 clientsecret是在github注册应用后获得的

const (
	state        = "mysite"
	clientid     = "41a7d0c279addd78477a"
	clientsecret = "e66f5e9e477d7e0097dde79a3416b862b7ffcc57"
	redirectURI  = "http://localhost:8080/authorization_code"
)

// Github ...
type Github struct {
	code      string
	scope     string
	token     string
	tokenType string
}

// New ...
func New() *Github {
	return &Github{}
}

// SetCode ...
func (g *Github) SetCode(code string) {
	g.code = code
}

// GetToken ...
func (g *Github) GetToken() (err error) {
	url := "https://github.com/login/oauth/access_token"

	param := map[string]string{
		"client_id":     clientid,
		"client_secret": clientsecret,
		"code":          g.code,
		"state":         state,
		"redirect_uri":  redirectURI,
	}
	b, err := json.Marshal(param)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	tmp := struct {
		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token_type"`
	}{}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return
	}

	g.token = tmp.AccessToken
	g.scope = tmp.Scope
	g.tokenType = tmp.TokenType
	return
}

// GetUser ...
func (g *Github) GetUser() (name string, err error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/user", nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "token "+g.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("用户信息：", string(b))

	tmp := struct {
		Login string
	}{}

	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return
	}
	name = tmp.Login

	// store user info to database
	return
}
