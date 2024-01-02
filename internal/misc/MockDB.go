package misc

import "sync"

type mockDB struct {
	UserURIs []string
}

var singleton *mockDB
var once sync.Once

func GetDB() *mockDB {
	once.Do(func() {
		singleton = &mockDB{}
		singleton.UserURIs = []string{}
	})
	return singleton
}
