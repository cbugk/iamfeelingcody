package model

import (
)

type GithubUser interface {
	Name() string
	setName(string)
	URI() GithubUserURI
	setURI(GithubUserURI)
}

type githubUser struct {
	name string
	uri GithubUserURI
}

func (user githubUser) Name() string {
	return user.name
}
func (user githubUser) setName(name string) {
	user.name = name
}

func (user githubUser) URI() GithubUserURI {
	return user.uri
}
func (user githubUser) setURI(uri GithubUserURI) {
	user.uri = uri
}

func MakeGithubUser(name string) (user GithubUser) {
	user = new(githubUser)
	user.setName(name)
	user.setURI(MakeGithubUserURI(name))
	return
}