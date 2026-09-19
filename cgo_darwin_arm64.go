//go:build darwin && arm64

package dengjen

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/darwin_arm64 -llibdengjen -Wl,-rpath,${SRCDIR}/lib/darwin_arm64
*/
import "C"
