package github

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/cbugk/iamfeelingcody/src/internal/sqlc"
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc/sqlite"
)

func PutUser(name string) (ok bool) {
	var user sqlite.GithubUser
	var errDB error
	var errGithub error
	var shouldUpdate bool

	log.Println("PutUser: ", name)

	// Name not provided
	if len(name) == 0 {
		log.Println("Name empty")
		return false
	}

	// Fetch DB entry
	user, errDB = sqlc.Q().GetGithubUser(context.Background(), name)
	// User already created
	if errDB == nil {
		log.Println(user.Name)
		// Hit 24-hour cache
		if user.Timestamp.Valid && (time.Now().Sub(user.Timestamp.Time)) < (24*time.Hour) {
			log.Println("Found existing user (", user.Timestamp.Time, ")")
			return true
		}
		// Miss 24-hour cache
		shouldUpdate = true
	}
	errGithub = CheckGithubUser(name)

	// Github user's url exists
	if errGithub == nil {
		if shouldUpdate {
			user, errDB = sqlc.Q().UpdateGithubUser(context.Background(), sqlite.UpdateGithubUserParams{
				Name:    name,
				Present: true,
			})
		} else {
			user, errDB = sqlc.Q().CreateGithubUser(context.Background(), sqlite.CreateGithubUserParams{
				Name:    name,
				Present: true,
			})
		}

		if errDB != nil {
			log.Println(errDB.Error(), name)
			return false
		}
		log.Println(user.Name)
		return true
	}

	// Github user's url does not exist
	if errors.Is(errGithub, &ErrorUserNotFound{}) {
		log.Println(errGithub.Error(), name)
		return false
	}

	// Github rate limit reached
	if errors.Is(errGithub, &ErrorTooManyRequests{}) {
		log.Println(errGithub.Error(), name)
		log.Println("Sleeping for 5 seconds.")
		time.Sleep(5 * time.Second)
		return false
	}

	// Unanticipated error
	log.Fatalln(errGithub.Error(), name)
	return false
}
