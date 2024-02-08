package check

type ErrorGithubTooManyRequests struct{}

func (e *ErrorGithubTooManyRequests) Error() string {
	return "Error: Github too many requests."
}
