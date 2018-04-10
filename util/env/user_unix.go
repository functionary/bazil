// +build darwin freebsd linux netbsd openbsd

package env // import "bazil.org/bazil/util/env"

import (
	"syscall"
)

func init() {
	MyUID = uint32(syscall.Getuid())
	MyGID = uint32(syscall.Getgid())
}
