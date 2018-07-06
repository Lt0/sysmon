// The dev pseudo-file contains network device status information.  This gives the number of received and sent pack‐
// ets,  the  number of errors and collisions and other basic statistics.  These are used by the ifconfig(8) program
// to report device status
package net

type Dev struct {
	Interface	[]string
}