# GoResolveHosts

This Golang script is used to convert a hostname to an IP address (or several IP addresses).

## Examples
To get an IP address for hostnames in file `hostnames.txt`: 
```
$ cat /tmp/hosts.txt
www.google.com
www.gmail.com

$ cat /tmp/hosts.txt | go run goresolvehosts.go
www.gmail.com|142.250.66.197
www.google.com|142.250.66.164
```


To get ALL IP address for hostnames in file `hostnames.txt`: 
```
$ cat /tmp/hosts.txt
www.google.com
www.gmail.com

$ cat /tmp/hosts.txt | go run goresolvehosts.go -allIPs
www.gmail.com|142.250.66.197,2404:6800:4006:80f::2005
www.google.com|142.250.66.164,2404:6800:4006:80e::2004
```