# To Do

- [x] map for guid->ipaddress lookups
- [x] hostname normalization
- [ ] thread-safety for map
- [x] /debug.txt - list all entries in map
- [x] /api.json
- [x] parameter for base hostname
- [x] parameter for public ip address
- [ ] https support
- [x] /status.json
- [x] /favicon.ico/svg
- [x] /robots.txt
- [ ] logging
- [ ] better root text (only on subdomains)
- [ ] root redirect to subdomain

- [ ] deployable pkg (apt? snap?)
- [ ] local resolver should not be 127.0.0.1

- [ ] support for dns lookup of version.bind, hostname.bind
- [ ] use redis instead of local map

## Resources

- [certmagic](https://github.com/mholt/certmagic)
- [example DNSProvider](https://github.com/go-acme/lego/blob/master/providers/dns/acmedns/acmedns.go) for certmagic
- [miekg reflect.go sample](https://github.com/miekg/exdns/blob/master/reflect/reflect.go)
- [go maps](https://blog.golang.org/go-maps-in-action)

## Ubuntu server DNS setup
```
systemctl disable systemd-resolved.service
```