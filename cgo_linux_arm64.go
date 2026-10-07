//go:build linux && arm64

package dengjen

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/linux_arm64 -llibdengjen -Wl,-rpath,${SRCDIR}/lib/linux_arm64
*/
import "C"
