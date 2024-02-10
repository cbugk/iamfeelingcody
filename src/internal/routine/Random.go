package routine

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/cbugk/iamfeelingcody/src/internal/alpnumd"
	"github.com/cbugk/iamfeelingcody/src/internal/github"
	pkgRoutine "github.com/cbugk/iamfeelingcody/src/pkg/routine"
	"github.com/dchest/uniuri"
)

func Random(n int, found chan<- string) {
	var wg sync.WaitGroup
	wg.Add(n)

	// Spawn routines for size of the request
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			// Until stop signal recieved
			// OR user is put into DB
			for !pkgRoutine.IsControlCRecieved {
				random := fmt.Sprint(uniuri.NewLenChars(1, alpnumd.Alpnum), uniuri.NewLenChars(rand.Intn(github.MaxLength-1), alpnumd.Alpnumdash))
				if github.PutUser(random) {
					log.Println("Has Put: ", random)
					found <- random
					break
				}
			}
		}()
	}

	wg.Wait()
	close(found)
}
