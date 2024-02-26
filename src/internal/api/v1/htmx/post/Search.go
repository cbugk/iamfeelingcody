package post

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc/sqlite"
	"github.com/julienschmidt/httprouter"
)

func Search(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "text/html")

	r.ParseForm()

	var users []sqlite.GithubUser
	var err error
	search := r.Form.Get("search")
	fmt.Println("Received search: ", search, r.Form.Encode())

	if len(search) != 0 {
		users, err = sqlc.Q().SearchLikeGithubUsers(ctx, sql.NullString{String: search, Valid: true})
	} else {
		users, err = sqlc.Q().ListGithubUsers(ctx)
	}

	// DB search error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	var sb strings.Builder
	for _, user := range users {
		sb.WriteString(fmt.Sprintf("<tbody><tr><th>%v</th><th>%v</th></tr></tbody>", user.Name, user.Timestamp.Time.Format("2006-01-02T15:04:05")))
	}
	w.Write([]byte(sb.String()))
}
