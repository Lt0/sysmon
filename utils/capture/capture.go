package capture

import (
	"fmt"
	"net"
	"github.com/google/gopacket"
	"github.com/google/gopacket/afpacket"
	"github.com/google/gopacket/layers"
	// "github.com/google/gopacket/pcap"
)

func CaptureDev(iface net.Interface) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic:", e)
			CaptureDev(iface)
		}
	}()

	devName := iface.Name
	HWAddr := iface.HardwareAddr.String()
	fmt.Println("capturing", devName)
	source, err := afpacket.NewTPacket(
		afpacket.OptInterface(devName),
		afpacket.OptBlockSize(1<<15))

	if err != nil {
		fmt.Printf("new t packet on dev %v failed\n", devName)
		return
	}

	defer source.Close()

	

	// map[pid]bytes
	// map[pid]packetsNum
	// var BytesRecord map[string]uint64
	// var PacketsRecord map[string]uint64

	// bytes := uint64(0)
	// packets := uint64(0)
	for {
		data, _, err := source.ZeroCopyReadPacketData()
		if err != nil {
				continue
		}

		packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)

		srcIP, dstIP := packet.NetworkLayer().NetworkFlow().Endpoints()
		srcPort, dstPort := packet.TransportLayer().TransportFlow().Endpoints()
		srcLink, dstLink := packet.LinkLayer().LinkFlow().Endpoints()
		fmt.Printf("%v - %v:%v -> %v - %v:%v\n", srcLink, srcIP, srcPort, dstLink, dstIP, dstPort)
		if srcLink.String() == HWAddr {
			fmt.Println("counting src port", srcPort)
			
		} else if dstLink.String() == HWAddr {
			fmt.Println("counting dst port", dstPort)
		} else {
			fmt.Println("counting forwarding")
		}


// 		bytes += uint64(len(data))
// 		packets++
	}
}

// func port2pid(port string) string {

// }

// func syncPids(pids map) {

// }

func capture() {
	source, err := afpacket.NewTPacket(
			afpacket.OptInterface("ens3"),
			afpacket.OptBlockSize(1<<20))

	if err != nil {
			fmt.Println("new t packet failed:", err)
	}

	defer source.Close()

	bytes := uint64(0)
	packets := uint64(0)

	for {
			data, _, err := source.ZeroCopyReadPacketData()
			if err != nil {
					fmt.Println("read packet err:", err)
					continue
			}
			bytes += uint64(len(data))
			packets++
//              fmt.Printf("read %d bytes in %d packets\n", bytes, packets)
			packet := gopacket.NewPacket(data, layers.LayerTypeEthernet, gopacket.NoCopy)
			tp := packet.TransportLayer()
			//fmt.Println("tp:", tp)
			flow := tp.TransportFlow()
			src, dst := flow.Endpoints()
			//fmt.Println("src port:", src)
			//fmt.Println("dst port:", dst)

			net_src, net_dst := packet.NetworkLayer().NetworkFlow().Endpoints()
			//fmt.Println("net_src:", net_src)
			//fmt.Println("net_dst:", net_dst)
			fmt.Printf("%v:%v -> %v:%v | size: %v\n", net_src, src, net_dst, dst, bytes)
	}
}
