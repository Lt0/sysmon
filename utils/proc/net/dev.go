// The dev pseudo-file contains network device status information.  This gives the number of received and sent pack‐
// ets,  the  number of errors and collisions and other basic statistics.  These are used by the ifconfig(8) program
// to report device status
package net

import (
	"strconv"
	"fmt"
	"strings"
	"bufio"
	"os"
	"io"
	"path/filepath"
)

type Dev struct {
	Ifaces		[]Iface
}

type Iface struct {
	Name		string
	Rx			RX
	Tx			TX
}

type RX struct {
	Bytes		uint64
	Packets		uint64
	Errs		uint64
	Drop		uint64
	Fifo		uint64
	Frame		uint64
	Compressed	uint64
	Multicast	uint64
}

type TX struct {
	Bytes		uint64
	Packets		uint64
	Errs		uint64
	Drop		uint64
	Fifo		uint64
	Colls		uint64
	Carrier		uint64
	Compressed	uint64
}

func (p *Dev) Update() error {
	f, err := os.Open(filepath.Join(Ctx.Procfs, netfs, "dev"))
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	r.ReadBytes('\n')
	_, err = r.ReadBytes('\n')
	if err == io.EOF {
		return fmt.Errorf("Invalid Ctx.Procfs file %v/%v/dev\n", Ctx.Procfs, netfs)
	}

	p.Ifaces = make([]Iface, 0)
	for {
		b, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		d, err := parseDev(string(b))
		if err != nil {
			continue
		}
		p.Ifaces = append(p.Ifaces, d)
	}

	return nil
}

func (p *Dev) CountTx() uint64 {
	var c uint64
	for _, v := range(p.Ifaces) {
		c += v.Rx.Bytes
	}

	return c
}

func (p *Dev) CountRx() uint64 {
	var c uint64
	for _, v := range(p.Ifaces) {
		c += v.Tx.Bytes
	}

	return c
}

func parseDev(s string) (Iface, error) {
	var iface Iface
	s = strings.TrimSpace(s)
	//type and value
	tv := strings.Split(s, ":")
	if len(tv) < 2 {
		return iface, fmt.Errorf("Invalid net dev info line: %v\n", s)
	}
	iface.Name = tv[0]
	
	vs := strings.Fields(tv[1])
	if len(vs) < 16 {
		return iface, fmt.Errorf("Invalid net dev info line: %v\n", s)
	}

	iface.Rx.Bytes, _ = strconv.ParseUint(vs[0], 10, 64)
	iface.Rx.Packets, _ = strconv.ParseUint(vs[1], 10, 64)
	iface.Rx.Errs, _ = strconv.ParseUint(vs[2], 10, 64)
	iface.Rx.Drop, _ = strconv.ParseUint(vs[3], 10, 64)
	iface.Rx.Fifo, _ = strconv.ParseUint(vs[4], 10, 64)
	iface.Rx.Frame, _ = strconv.ParseUint(vs[5], 10, 64)
	iface.Rx.Compressed, _ = strconv.ParseUint(vs[6], 10, 64)
	iface.Rx.Multicast, _ = strconv.ParseUint(vs[7], 10, 64)

	iface.Tx.Bytes, _ = strconv.ParseUint(vs[8], 10, 64)
	iface.Tx.Packets, _ = strconv.ParseUint(vs[9], 10, 64)
	iface.Tx.Errs, _ = strconv.ParseUint(vs[10], 10, 64)
	iface.Tx.Drop, _ = strconv.ParseUint(vs[11], 10, 64)
	iface.Tx.Fifo, _ = strconv.ParseUint(vs[12], 10, 64)
	iface.Tx.Colls, _ = strconv.ParseUint(vs[13], 10, 64)
	iface.Tx.Carrier, _ = strconv.ParseUint(vs[14], 10, 64)
	iface.Tx.Compressed, _ = strconv.ParseUint(vs[15], 10, 64)

	return iface, nil
}