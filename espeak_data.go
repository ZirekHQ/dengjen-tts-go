package dengjen

import "os"

// init sets DENGJEN_ESPEAKNG_DATA_DIRECTORY to this module's bundled
// espeak-ng-data directory, unless the environment already has it set —
// so a fresh `go get` works with no manual setup, while an operator who
// has already pointed it elsewhere is never overridden. When the bundled
// directory can't be located (see defaultEspeakDataDir), the variable is
// left alone rather than set to a bogus path.
func init() {
	if _, set := os.LookupEnv("DENGJEN_ESPEAKNG_DATA_DIRECTORY"); set {
		return
	}
	if dir := defaultEspeakDataDir(); dir != "" {
		_ = os.Setenv("DENGJEN_ESPEAKNG_DATA_DIRECTORY", dir)
	}
}
