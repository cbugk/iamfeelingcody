package embed

import (
	"embed"
	"io/fs"
)

//go:embed public
var public embed.FS

func Dir() fs.FS {
	sub, err := fs.Sub(public, "public")
	if err != nil {
		panic(err)
	}
	return sub
}
