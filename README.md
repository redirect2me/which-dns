# DNS Resolver Detection

This is a combined DNS+HTTP server that shows which DNS server is being used.


## How does it work?

1. You make a request to a hostname with a unique prefix.
2. All hostnames resolve to the same IP, but the DNS server records which IP address the query came from.
3. The webserver looks for this record and returns it.

