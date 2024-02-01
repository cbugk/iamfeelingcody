package embed

import (
	"embed"
	"io/fs"
)

//go:embed static
var static embed.FS

func Dir() fs.FS {
	sub, err := fs.Sub(static, "static")
	if err != nil {
		panic(err)
	}
	return sub
}
