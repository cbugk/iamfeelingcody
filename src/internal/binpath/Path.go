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

func Base() string {
	_, base := Path()
	return base
}

func Dir() string {
	dir, _ := Path()
	return dir
}

func PublicDir() string {
	return filepath.Join(Dir(), "public")
}
