package check

import (
	"fmt"
	"log"
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
	} else if resp.StatusCode == 429 {
		return &ErrorGithubTooManyRequests{}
	} else if resp.StatusCode != 404 {
		log.Fatalln(resp.StatusCode)
		return nil
	} else {
		return &ErrorGithubUserNotFound{}
	}
}
