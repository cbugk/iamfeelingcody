package misc

import (
	"fmt"
	"net/http"

	"github.com/dchest/uniuri"
)

func FindGithubUser(length int) string {
	userURI := ""

	for true {
		userURI = fmt.Sprintf("https://github.com/%s", uniuri.NewLen(length))

		if resp, err := http.Get(userURI); err != nil {
			fmt.Println(err)
		} else if resp.StatusCode != 200 {
			fmt.Printf("Non-OK status code: %d, %s\n", resp.StatusCode, userURI)
		} else {
			fmt.Printf("Found: %s\n", userURI)
			break
		}
	}

	return userURI
}
