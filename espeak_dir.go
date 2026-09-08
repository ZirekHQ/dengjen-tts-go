package dengjen

import (
	"os"
	"path/filepath"
	"runtime"
)

func defaultEspeakDataDir() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}
	dir := filepath.Join(filepath.Dir(file), "espeak-ng-data")
	if !filepath.IsAbs(dir) {
		return ""
	}
	if _, err := os.Stat(dir); err != nil {
		return ""
	}
	return dir
}
