package github

type ErrorUserNotFound struct{}

func (e *ErrorUserNotFound) Error() string {
	return "Error: Github user not found."
}
