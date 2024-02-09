package routine

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	pkgRoutine "github.com/cbugk/iamfeelingcody/src/pkg/routine"
	"github.com/cbugk/iamfeelingcody/src/internal/github"
	"github.com/cbugk/iamfeelingcody/src/internal/ralpv"
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
			for !pkgRoutine.ShouldStop {
				random := fmt.Sprint(uniuri.NewLenChars(1, ralpv.Alpnum), uniuri.NewLenChars(rand.Intn(github.MaxLength-1), ralpv.Alpnumdash))
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
