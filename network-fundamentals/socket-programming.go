/* IP */
package main

import (
	"fmt"
	"net"
	"os"
)

func validateIPAddress() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1) }
		name := os.Args[1]
		addr := net.ParseIP(name) 
		
		if addr == nil {
		fmt.Println("Invalid address") } else {
		fmt.Println("The address is ", addr.String()) }
		os.Exit(0)
}

func IPMask() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1) }
		dotAddr := os.Args[1]
		addr := net.ParseIP(dotAddr)
		 if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1) }
		mask := addr.DefaultMask()
		network := addr.Mask(mask)
		ones, bits := mask.Size()
		 fmt.Println("Address is ", addr.String(),
		" Default mask length is ", bits, "Leading ones count is ", ones, "Mask is (hex) ", mask.String(), " Network is ", network.String())
		os.Exit(0) 
}

func ResolveIPAddr() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
		}
		name := os.Args[1]
		addr, err := net.ResolveIPAddr("ip", name) 
		if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1) }
		fmt.Println("Resolved address is ", addr.String())
		os.Exit(0)
}

func HostLookUp() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1) }
		name := os.Args[1]
		addrs, err := net.LookupHost(name)
		 if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2) }
		for _, s := range addrs { fmt.Println(s)
		}
		os.Exit(0) 
}

func LookUpPort() {

		if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
		"Usage: %s network-type service\n", os.Args[0])
		os.Exit(1) }
		networkType := os.Args[1] 
		service := os.Args[2]
		port, err := net.LookupPort(networkType, service) 
		if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2) }
		fmt.Println("Service port ", port)
		os.Exit(0) 
}

func main() {
	if len(os.Args) != 3 {
        fmt.Fprintf(os.Stderr, "Usage: %s network-address port\n", os.Args[0])
        os.Exit(1)
    }

    networkAddress := os.Args[1]
    port := os.Args[2]

    tcpAddr, err := net.ResolveTCPAddr("tcp", networkAddress+":"+port)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
        os.Exit(2)
    }

    fmt.Printf("Resolved address: %s\n", tcpAddr.String())
    os.Exit(0)
}
