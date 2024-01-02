package misc

import "fmt"

func GithubUserURI(u string) string {
	return fmt.Sprintf("https://github.com/%s", u)
}
