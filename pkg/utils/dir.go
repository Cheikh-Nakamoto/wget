package utils

import (
	"os"
	"path/filepath"
	"strings"
	"wget/pkg"
)

func SetOutputDir(dir string) {
	os.Mkdir(dir, 0755)
	*pkg.NewPath = dir
}

func AbsolutePath(dir string) (string, error) {
	if strings.HasPrefix(dir, "~/") {
		dir = strings.TrimPrefix(dir, "~/")
		home, err := os.UserHomeDir()
		return filepath.Join(home, dir), err
	} else if strings.HasPrefix(dir, "../") || strings.HasPrefix(dir, "./") {
		return dir, nil
	} else {
		// strings.TrimPrefix(dir, "~/")
		return filepath.Abs(filepath.Clean(filepath.Base(dir)))
	}
}
