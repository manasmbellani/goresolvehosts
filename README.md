# GoResolveHosts

This Golang script is used to convert a hostname to an IP address (or several IP addresses) OR perform reverse PTR on IP addresses to get a hostname.

## Examples
To get an IP address for hostnames in file `/tmp/hosts.txt`: 
```
$ cat /tmp/hosts.txt
www.google.com
www.gmail.com

$ cat /tmp/hosts.txt | go run goresolvehosts.go
www.gmail.com|142.250.66.197
www.google.com|142.250.66.164
```

To get ALL IP address for hostnames in file `/tmp/hosts.txt` with 50 go-routines ("light-threads") run the command with `-numThreads` flag: 
```
$ cat /tmp/hosts.txt
www.google.com
www.gmail.com

$ cat /tmp/hosts.txt | go run goresolvehosts.go -allIPs -numThreads 50
www.gmail.com|142.250.66.197,2404:6800:4006:80f::2005
www.google.com|142.250.66.164,2404:6800:4006:80e::2004
```

To perform a reverse PTR on a set of IP addresses to get the domain names, run the following command. Please note that non-resolvable hosts does not display a response.
```
$ cat /tmp/ips.txt 
1.1.1.1
127.0.0.1
7.22.11.1
142.250.66.196

$ cat /tmp/ip.txt | go run goresolvehosts.go -r
127.0.0.1|localhost
1.1.1.1|one.one.one.one.
142.250.66.196|syd09s23-in-f4.1e100.net.
```

To get the response only (and not display the query) for the A/PTR record, run the command:
```
$ cat /tmp/ips.txt | go run goresolvehosts.go -r -respOnly
localhost
one.one.one.one.
syd09s23-in-f4.1e100.net.
```