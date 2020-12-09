package http

import (
	"fmt"
	"net/http"

	"site/oauth"
	"site/oauth/github"
)

var (
	githubAuth oauth.OAuth
)

func init() {
	githubAuth = github.New()

	http.HandleFunc("/authorization_code", authorizationCode)
}

// http://localhost:8080/authorization_code?code=3e4c4be65019ce587c02&state=site
func authorizationCode(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad code"))
		return
	}

	githubAuth.SetCode(code)
	fmt.Println("已得到授权码！")

	err := githubAuth.GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Token获取成功！")

	name, err := githubAuth.GetUser()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("用户信息获取成功！")

	w.WriteHeader(200)
	w.Write([]byte("已认证成功，欢迎【" + name + "】！"))
	return
}
