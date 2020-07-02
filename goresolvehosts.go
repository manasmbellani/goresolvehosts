package main

// Golang Script to resolve hosts to an IP address and write it to output

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// Delim - Delimiter to use when writing Hostname to IP address
const Delim = "|"

// DNSTimeout - Timeout to use for Reverse PTR
const DNSTimeout = 1000 * time.Millisecond

func main() {
	allIPPtr := flag.Bool("allIPs", false,
		"Set this flag to show ALL IPs, not just one")
	returnRespOnlyPtr := flag.Bool("respOnly", false,
		"Set this flag to return the response only - not the query")
	numThreadsPtr := flag.Int("numThreads", 20,
		"Get the number of threads to use")
	reversePtr := flag.Bool("r", false, "Perform reverse PTR to get hostnames for IP addresses ")
	flag.Parse()
	allIPs := *allIPPtr
	returnRespOnly := *returnRespOnlyPtr
	numThreads := *numThreadsPtr
	reverse := *reversePtr

	// List of hosts to resolve
	hosts := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Wait to receive the next host
			for host := range hosts {

				if reverse {
					// need to perform Reverse PTR lookup on domain name
					ctx, cancel := context.WithTimeout(context.TODO(),
						DNSTimeout)
					defer cancel()

					// Get the host names from the IP address
					var r net.Resolver
					names, err := r.LookupAddr(ctx, host)
					if err == nil && len(names) > 0 {

						if returnRespOnly {
							// only display response
							fmt.Println(names[0])
						} else {
							// Display the query and response
							fmt.Printf("%s%s%s\n", host, Delim, names[0])
						}
					}
				} else {
					ips, _ := net.LookupIP(host)

					if ips != nil {
						ipsToShow := ""
						// Convert all IPs to string
						var ipsAsStr []string

						// Print the hostname only
						if returnRespOnly {
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
				}
			}
		}()
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		host := sc.Text()

		if host != "" {
			// Add the hosts to resolve from the user input if not null
			hosts <- host
		}
	}

	// no more hosts to send
	close(hosts)

	wg.Wait()
}
