package dengjen

import "os"

func init() {
	if _, set := os.LookupEnv("DENGJEN_ESPEAKNG_DATA_DIRECTORY"); set {
		return
	}
	if dir := defaultEspeakDataDir(); dir != "" {
		_ = os.Setenv("DENGJEN_ESPEAKNG_DATA_DIRECTORY", dir)
	}
}
