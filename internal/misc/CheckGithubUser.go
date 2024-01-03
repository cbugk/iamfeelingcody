package misc

import (
	"net/http"
)

func CheckGithubUser(u string) (string, error) {
	userURI := GithubUserURI(u)

	resp, err := http.Get(userURI)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return userURI, nil
	} else {
		return userURI, &ErrorGithubUserNotFound{}
	}
}
