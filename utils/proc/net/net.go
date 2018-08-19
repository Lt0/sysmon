// various net pseudo-files, all of which give the status of some part of the networking layer.  
// These files contain ASCII structures and are, therefore, readable with cat(1). 
// However, the standard netstat(8) suite provides  much cleaner access to these files.
package net


var netfs = "net"

type Context struct {
	Procfs string
}
var Ctx Context

func init() {
	Ctx.Procfs = "/proc"
}
