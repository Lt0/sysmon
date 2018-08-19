// The proc filesystem is a pseudo-filesystem which provides an interface to kernel data structures.  
// It is commonly mounted at /proc. Most of it is read-only, but some files allow kernel variables to be changed.
package proc

import (
	//"fmt"
)

type Context struct {
	Procfs string
}
var Ctx Context

func init() {
	Ctx.Procfs = "/proc"
}
