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
	flag.Parse()
	allIPs := *allIPPtr

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
					if allIPs {
						for _, ip := range ips {
							ipsAsStr = append(ipsAsStr, ip.String())
						}
						ipsToShow = strings.Join(ipsAsStr, ",")
					} else {
						ipsToShow = ips[0].String()
					}
					fmt.Printf("%s%s%s\n", host, Delim, ipsToShow)
				}
			}(host)
		}
	}
	wg.Wait()
}
