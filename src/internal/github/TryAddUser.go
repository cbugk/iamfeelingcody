package github

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/cbugk/iamfeelingcody/src/internal/ralpv"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc/sqlite"
)

func TryAddUser(name string) {
	for {
		log.Println("TryAddUser: ", name)

		if len(name) == 0 {
			// Name not provided
			log.Println("Name empty")
		} else if user, err := sqlc.Q().GetGithubUser(context.Background(), name); err == nil {
			// User already created
			log.Println(user.Ralpv, user.Name)
		} else if err := CheckGithubUser(name); err == nil {
			// Github user's url exists
			if user, err = sqlc.Q().CreateGithubUser(context.Background(), sqlite.CreateGithubUserParams{name, ralpv.NameToRalpv(name), true}); err != nil {
				log.Println(err.Error())
			} else {
				log.Println(user.Ralpv, user.Name)
			}
		} else if errors.Is(err, &ErrorUserNotFound{}) {
			// Github user's url does not exist
			log.Println(err.Error())
			if user, err = sqlc.Q().CreateGithubUser(context.Background(), sqlite.CreateGithubUserParams{name, ralpv.NameToRalpv(name), false}); err != nil {
				log.Println(err.Error())
			} else {
				log.Println(user.Ralpv, user.Name)
			}
		} else if errors.Is(err, &ErrorTooManyRequests{}) {
			log.Println(err.Error())
			log.Println("Sleeping for 5 seconds.")
			time.Sleep(5 * time.Second)
		} else {
			// Unanticipated error
			log.Fatalln(err.Error())
		}
	}
}
