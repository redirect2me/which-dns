# Which DNS: DNS Resolver Detection [<img alt="Which DNS Logo" src="assets/favicon.svg" height="90" align="right"/>](https://resolve.rs/)

This is a combined DNS+HTTP server that shows which DNS server is being used.


## How does it work?

1. You make a request to a hostname with a unique prefix.
2. All hostnames resolve to the same IP, but the DNS server records which IP address the query came from.
3. The webserver looks for this record and returns it.

## License

[GNU Affero General Public License v3.0](LICENSE.txt)

## Credits

[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![Google Analytics](https://www.vectorlogo.zone/logos/google_analytics/google_analytics-ar21.svg)](https://www.google.com/analytics "Traffic Measurement")
[![Ubuntu](https://www.vectorlogo.zone/logos/ubuntu/ubuntu-ar21.svg)](https://www.ubuntu.com/ "Word list")
