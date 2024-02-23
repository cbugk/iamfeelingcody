package htmx

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"github.com/cbugk/iamfeelingcody/src/internal/routine"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc/sqlite"
	"github.com/cbugk/iamfeelingcody/src/internal/templ/page"
	"github.com/julienschmidt/httprouter"
)

func Random(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	n := 0
	var err error
	users := make([]sqlite.GithubUser, 0)

	s := r.URL.Query().Get("n")
	log.Printf("s: '%v'\n", s)
	if regexp.MustCompile(`[0-9]*`).MatchString(s) {
		n, err = strconv.Atoi(s)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		n = rand.Intn(5) + 1 // zero is not valid
		log.Println("Random n not provided, generated: ", n)
	}

	if n > 0 {
		found := make(chan string, n)
		routine.Random(n, found)
		// Accumulate results
		for f := range found {
			if user, err := sqlc.Q().GetGithubUser(context.Background(), f); err != nil {
				log.Println(err.Error())
			} else {
				users = append(users, user)
			}
		}
		// found closed by Random
	}

	page.Random(users).Render(r.Context(), w)
}
