package misc

type ErrorGithubUserNotFound struct{}

func (e *ErrorGithubUserNotFound) Error() string {
	return "Error: Github user not found."
}
