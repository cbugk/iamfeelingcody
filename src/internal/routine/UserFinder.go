package routine

import (
	"fmt"
	"math/rand"

	"github.com/cbugk/iamfeelingcody/src/internal/alpnumd"
	"github.com/cbugk/iamfeelingcody/src/internal/github"
	"github.com/dchest/uniuri"
)

func UserFinder(names <-chan string, n int) func() {
	return func() {
		for i := 0; i < n; i++ {
			// Until requested number of users are put into DB
			for {
				if n <= 0 || github.PutUser(fmt.Sprint(uniuri.NewLenChars(1, alpnumd.Alpnum), uniuri.NewLenChars(rand.Intn(github.MaxLength-1), alpnumd.Alpnumdash))) {
					break
				}
			}
		}
	}
}
