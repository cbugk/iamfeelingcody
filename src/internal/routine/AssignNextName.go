package routine

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/cbugk/iamfeelingcody/src/internal/ralpv"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
)

func AssignNextName(names chan<- string) {
	go func() {
		for {
			user, err := sqlc.Q().HighestRalpvUser(context.Background())
			if err != nil {
				if errors.As(err, &sql.ErrNoRows) {
					names <- ralpv.RalpvToName(0)
				} else {
					// Unknown error
					log.Println(err.Error())
				}
			} else {
				names <- ralpv.RalpvToName(user.Alph + 1)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}
