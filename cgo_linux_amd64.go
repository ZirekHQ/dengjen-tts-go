//go:build linux && amd64

package dengjen

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/linux_amd64 -llibdengjen -Wl,-rpath,${SRCDIR}/lib/linux_amd64
*/
import "C"
