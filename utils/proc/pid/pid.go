// There is a numerical subdirectory for each running process; the subdirectory is named by the process ID.  
// Each such subdirectory contains the following pseudo-files and directories.
package pid

type Context struct {
	Procfs string	// proc filesystem mount point, default is /proc
}
var Ctx Context

func init() {
	Ctx.Procfs = "/proc"
}