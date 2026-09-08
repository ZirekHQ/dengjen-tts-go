//go:build windows && amd64

package dengjen

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib/windows_amd64 -llibdengjen
*/
import "C"
