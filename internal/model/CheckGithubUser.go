package model

import (
	"fmt"
	"net/http"
)

func CheckGithubUser(name string) error {
	userURI := fmt.Sprintf("https://github.com/%v", name)

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
