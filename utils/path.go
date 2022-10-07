package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetProjectPath get project work dir
func GetProjectPath() string {
	workDir, _ := os.Getwd()
	path := workDir
	count := 0
	maxDeepth := 32
	for {
		if count > maxDeepth {
			panic("reach max deepth")
		}

		if path == "/" {
			panic(fmt.Sprintf("can't find assets dir base on %s", workDir))
		}

		assetsDir := filepath.Join(path, "assets")

		if file, err := os.Stat(assetsDir); err == nil && file.IsDir() {
			return path
		}

		path = filepath.Join(path, "../")

		count++
	}
}
