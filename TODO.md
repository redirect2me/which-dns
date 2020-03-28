# To Do

## Coding

- [ ] [LRU cache](https://github.com/hashicorp/golang-lru) instead of map
- [ ] http request logging
- [ ] better root text (only on subdomains)
- [ ] debug page to json
- [ ] another debug page with distinct DNS servers
- [ ] root redirect to subdomain

## Packaging

- [ ] init script
- [ ] deployable pkg (apt? snap? terraform?)
- [ ] local resolver should not be 127.0.0.1

## Future

- [ ] test IPv6
- [ ] support for dns lookup of version.bind, hostname.bind
- [ ] use redis instead of local map
- [ ] [self updating](https://github.com/inconshreveable/go-update)?

## Done

- [x] map for guid->ipaddress lookups
- [x] hostname normalization
- [x] /debug.txt - list all entries in map
- [x] /api.json
- [x] parameter for base hostname
- [x] parameter for public ip address
- [x] https support
- [x] /status.json
- [x] /favicon.ico/svg
- [x] /robots.txt

## Resources

- [certmagic](https://github.com/mholt/certmagic)
- [example DNSProvider](https://github.com/go-acme/lego/blob/master/providers/dns/acmedns/acmedns.go) for certmagic
- [miekg reflect.go sample](https://github.com/miekg/exdns/blob/master/reflect/reflect.go)
- [go maps](https://blog.golang.org/go-maps-in-action)
- [acme-dns](https://github.com/joohoi/acme-dns)

## Ubuntu server DNS setup
```
systemctl disable systemd-resolved.service
```

and in `/etc/resolv.conf`

```
nameserver 8.8.8.8
```
