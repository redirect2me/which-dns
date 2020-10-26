# Which DNS: DNS Resolver Detection [<img alt="Which DNS Logo" src="assets/favicon.svg" height="90" align="right"/>](https://which.nameserve.rs/)

This is a combined DNS+HTTP server that shows which DNS server a computer is using.


## How does it work?

1. You make a request to a hostname with a unique prefix.
2. All hostnames resolve to the same IP, but the DNS server records which IP address the query came from.
3. The webserver looks for this record and returns it.

## Using

My server is running at `which.nameserve.rs`.

Be forewarned: it is running on the cheapest box I could find with a static IP.  You can hit it for light, non-commercial use.  I specifically made the API be JSONP only (i.e. you need to provide a `callback` parameter), so if you abuse it, bad things will happen to your clients!

Make an HTTPS request to `GUID.which.nameserve.rs/api.json?callback=myfunction`.  The `GUID` should be a unique string (not necessarily an actual GUID), different for every call. It will call `myfunction` with an object that has the following fields:

* `success` - boolean if it succeeded or not
* `output` - the result (if it succeeded)
* `message` - the error message (if it failed)

You can see it in action on the home page of [resolve.rs](https://resolve.rs/).

## Running your own copy

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

* email: the email address for your account with Let's Encrypt (and the SOA record)
* hostname: the `hostname` that you picked
* ipaddress: the public IP address of the server
* nshostname: the `nshostname` that you picked

## License

[GNU Affero General Public License v3.0](LICENSE.txt)

## Credits

[![certmagic](https://www.vectorlogo.zone/logos/github_mholt_certmagic/github_mholt_certmagic-ar21.svg)](https://github.com/mholt/certmagic "Certificate management")
[![Digital Ocean](https://www.vectorlogo.zone/logos/digitalocean/digitalocean-ar21.svg)](https://www.digitalocean.com/ "Hosting")
[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![GoatCounter](https://www.vectorlogo.zone/logos/goatcounter/goatcounter-ar21.svg)](https://www.goatcounter.com/ "Traffic Measurement")
[![Let's Encrypt](https://www.vectorlogo.zone/logos/letsencrypt/letsencrypt-ar21.svg)](https://letsencrypt.org/ "HTTPS certificates")
[![svgrepo](https://www.vectorlogo.zone/logos/svgrepo/svgrepo-ar21.svg)](https://www.svgrepo.com/svg/277712/witch "favicon (modified)")
[![Ubuntu](https://www.vectorlogo.zone/logos/ubuntu/ubuntu-ar21.svg)](https://www.ubuntu.com/ "Server operating system")
[![water.css](https://www.vectorlogo.zone/logos/netlifyapp_watercss/netlifyapp_watercss-ar21.svg)](https://watercss.netlify.app/ "Classless CSS")

## Alternatives

I wasn't the first person to come up with this idea.  Here are some other public sites that do the same thing:

* [dnsleaktest](https://dnsleaktest.com/)
* [whatsmydnsserver](http://www.whatsmydnsserver.com/)
* [ipleak.net](https://ipleak.net/) - a bit of information overload though.
* [benchmark from GRC](https://www.grc.com/dns/benchmark.htm) - fat-client tool to benchmark which DNS servers are fastest from your workstation.
* [dnsadblock](https://dnsadblock.com/dns-leak-test/)
* [akahelp.net](https://developer.akamai.com/blog/2018/05/10/introducing-new-whoami-tool-dns-resolver-information) - name server that you can use with `dig`