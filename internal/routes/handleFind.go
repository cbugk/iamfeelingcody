package routes

import (
	"fmt"
	"net/http"

	"github.com/cbugk/iamfeelingcody/internal/templates"
	"github.com/dchest/uniuri"
)

func findGithubUser(length int) string {
	userURI := ""

	for true {
		userURI = fmt.Sprintf("https://github.com/%s", uniuri.NewLen(length))

		if resp, err := http.Get(userURI); err != nil {
			fmt.Println(err)
		} else if resp.StatusCode != 200 {
			fmt.Println("Non-OK status code: %s, %s", resp.StatusCode, userURI)
		} else {
			fmt.Printf("Found: %s\n", userURI)
			break
		}
	}

	return userURI
}

func handleFind(w http.ResponseWriter, r *http.Request) {
	templates.PageFind(findGithubUser(1)).Render(r.Context(), w)
}
