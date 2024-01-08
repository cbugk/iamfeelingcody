package model

import (
	"fmt"
)

type GithubUserURI interface {
	Text() string
	setText(string)
}

type githubUserURI struct {
	text string
}

func(uri githubUserURI) Text() string {
	return uri.text
}
func(uri githubUserURI) setText(text string) {
	uri.text = text
}

func MakeGithubUserURI(name string) (uri GithubUserURI) {
	uri = new(githubUserURI)
	uri.setText(fmt.Sprintf("https://github.com/%s", name))
	return
}