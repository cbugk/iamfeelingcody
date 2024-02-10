package htmx

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/cbugk/iamfeelingcody/src/internal/github"
	"github.com/cbugk/iamfeelingcody/src/internal/routine"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc/sqlite"
	"github.com/cbugk/iamfeelingcody/src/internal/templ/page"
	"github.com/julienschmidt/httprouter"
)

func Random(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var n int
	var err error

	if s := r.URL.Query().Get("n"); s == "" {
		n = rand.Intn(github.MaxLength)
		log.Println("Random n not provided, generated: ", n)
	} else {
		n, err = strconv.Atoi(s)
		if err != nil {
			log.Println("Random n not valid, set to: ", n)
			n = 0
		}
	}

	found := make(chan string, n)
	routine.Random(n, found)
	// Accumulate results
	users := make([]sqlite.GithubUser, 0)
	for f := range found {
		if user, err := sqlc.Q().GetGithubUser(context.Background(), f); err != nil {
			log.Println(err.Error())
		} else {
			users = append(users, user)
		}
	}
	// found closed by Random

	page.Random(users).Render(r.Context(), w)
}
