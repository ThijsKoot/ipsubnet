package main

import (
	"flag"
	"fmt"

	"strconv"

	"strings"

	"os"

	"github.com/thijskoot/ipsubnet"
)

var (
	mask    int
	network string
)

type NetworkInfo struct {
	NumberOfIpaddresses      int      `json:"number-of-addresses"`
	NumberOfAddressableHosts int      `json:"number-of-available_addresses"`
	IpRange                  []string `json:"ip-range"`
	NetworkBits              int      `json:"network-bits"`
	BroadcastAddress         string   `json:"broadcast-address"`
}

func init() {
	flag.IntVar(&mask, "mask", 24, "set the subnet mask")
	flag.StringVar(&network, "network", "192.168.0.0", "network or host")
	flag.Parse()
}

func main() {
	sub := ipsubnet.SubnetCalculator(network, mask)

	if len(os.Args) == 2 {
		s, err := parseRange(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sub = s
	}

	n := NetworkInfo{
		NumberOfIpaddresses:      sub.GetNumberIPAddresses(),
		NumberOfAddressableHosts: sub.GetNumberAddressableHosts(),
		IpRange:                  sub.GetIPAddressRange(),
		NetworkBits:              sub.GetNetworkSize(),
		BroadcastAddress:         sub.GetBroadcastAddress(),
	}

	fmt.Printf("network:\t\t\t%s/%d\n", network, mask)
	fmt.Printf("number-of-addresses:\t\t%d\n", n.NumberOfIpaddresses)
	fmt.Printf("number-of-available-addresses:\t%d\n", n.NumberOfAddressableHosts)
	fmt.Printf("ip-range:\t\t\t[ %s - %s ]\n", n.IpRange[0], n.IpRange[1])
	fmt.Printf("broadcast-address:\t\t%s\n", n.BroadcastAddress)

}

func parseRange(s string) (*ipsubnet.Ip, error) {
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid input '%s'", s)
	}

	mask, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("unable to parse mask from input '%s'", parts[1])
	}

	return ipsubnet.SubnetCalculator(parts[0], mask), nil
}
