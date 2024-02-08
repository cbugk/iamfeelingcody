package github

type ErrorTooManyRequests struct{}

func (e *ErrorTooManyRequests) Error() string {
	return "Error: Github too many requests."
}
