package dengjen

import (
	"os"
	"path/filepath"
	"runtime"
)

// defaultEspeakDataDir returns the directory (bundled into this module)
// that holds espeak-ng-data, resolved via the source file's own location
// at build time. Under a -trimpath build (or any build where debug path
// info isn't embedded), runtime.Caller can't resolve a real path — in
// that case this returns "" and the caller (init, in espeak_data.go)
// leaves DENGJEN_ESPEAKNG_DATA_DIRECTORY unset rather than setting a
// bogus one, degrading to the same manual-configuration behavior the
// module's development copy already requires.
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
