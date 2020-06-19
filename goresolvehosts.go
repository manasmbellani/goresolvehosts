package main

// Golang Script to resolve hosts to an IP address and write it to output

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

// Delim - Delimiter to use when writing Hostname to IP address
const Delim = "|"

func main() {
	allIPPtr := flag.Bool("allIPs", false,
		"Set this flag to show ALL IPs, not just one")
	returnHostsOnlyPtr := flag.Bool("hostsOnly", false,
		"Set this flag to return host name only - not IP address")

	flag.Parse()
	allIPs := *allIPPtr
	returnHostsOnly := *returnHostsOnlyPtr

	var wg sync.WaitGroup

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		host := sc.Text()

		if host != "" {
			wg.Add(1)
			go func(host string) {
				defer wg.Done()
				ips, _ := net.LookupIP(host)

				if ips != nil {
					ipsToShow := ""
					// Convert all IPs to string
					var ipsAsStr []string

					// Print the hostname only
					if returnHostsOnly {
						fmt.Printf("%s\n", host)
					} else {
						if allIPs {
							// Show All Resolved IP
							for _, ip := range ips {
								ipsAsStr = append(ipsAsStr, ip.String())
							}
							ipsToShow = strings.Join(ipsAsStr, ",")
						} else {
							// Show the first IP from list of resolved IPs
							ipsToShow = ips[0].String()
						}
						fmt.Printf("%s%s%s\n", host, Delim, ipsToShow)
					}
				}
			}(host)
		}
	}
	wg.Wait()
}
