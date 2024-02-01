package binpath

import "path/filepath"

func PublicDir() string {
	return filepath.Join(Dir(), "public")
}
