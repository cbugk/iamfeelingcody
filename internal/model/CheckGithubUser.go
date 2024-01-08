package model

import (
	"net/http"
)

func CheckGithubUser(user GithubUser) error {
	userURI := user.URI().Text()

	resp, err := http.Get(userURI)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	} else {
		return &ErrorGithubUserNotFound{}
	}
}
