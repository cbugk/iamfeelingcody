package routine

import "github.com/cbugk/iamfeelingcody/src/internal/github"

func UserFinder(names <-chan string) func() {
	return func() {
		github.TryAddUser(<-names)
	}
}
