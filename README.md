# Which DNS: DNS Resolver Detection [<img alt="Which DNS Logo" src="assets/favicon.svg" height="90" align="right"/>](https://resolve.rs/)

This is a combined DNS+HTTP server that shows which DNS server a computer is using.


## How does it work?

1. You make a request to a hostname with a unique prefix.
2. All hostnames resolve to the same IP, but the DNS server records which IP address the query came from.
3. The webserver looks for this record and returns it.

## Running

You need a static IP address to run this server.  

Then you need to decide on two host names:

* `hostname` is the name of the subdomain that this server will own (for example: `which.example.com`)
* `nshostname` is the name of the authoritative name server for the subdomain (for example: `which-dns.example.com`)

They will both be served by the same server and at the same IP address.

The following DNS records need to be added to the main domain (for the above examples this would be `example.com`):

* A for the `nshostname` pointing to the static IP
* NS for the `hostname` pointing to the `nshostname`

The program needs access to the following ports:

* 53: for DNS
* 80: for HTTP
* 443: for HTTPS

The following parameters are required:

* email: the email address for your account with Let's Encrypt
* hostname: the `hostname` that you picked
* ipaddress: the public IP address of the server
* nshostname: the `nshostname` that you picked

## License

[GNU Affero General Public License v3.0](LICENSE.txt)

## Credits

[![Digital Ocean](https://www.vectorlogo.zone/logos/digitalocean/digitalocean-ar21.svg)](https://www.digitalocean.com/ "Hosting")
[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![Let's Encrypt](https://www.vectorlogo.zone/logos/letsencrypt/letsencrypt-ar21.svg)](https://letsencrypt.org/ "HTTPS certificates")
[![Ubuntu](https://www.vectorlogo.zone/logos/ubuntu/ubuntu-ar21.svg)](https://www.ubuntu.com/ "Server operating system")

* [certmagic](https://github.com/mholt/certmagic)
