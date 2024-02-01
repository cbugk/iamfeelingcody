package binpath

import (
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once
var base string
var dir string

func Path() (string, string) {
	once.Do(func() {
		e, err := os.Executable()
		if err != nil {
			panic(err)
		}

		dir = filepath.Dir(e)
		base = filepath.Base(e)
	})

	return dir, base
}
