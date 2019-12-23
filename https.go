package main

import (
	//"fmt"
	"github.com/go-acme/lego/v3/challenge/dns01"
	"github.com/mholt/certmagic"
)

type LocalDNSProvider struct {
}

func (d *LocalDNSProvider) Present(domain, token, keyAuth string) error {

	logger.Printf("INFO: presenting '%s'", domain)

	fqdn, value := dns01.GetRecord(domain, keyAuth)
	lookup_set("DNS01:"+fqdn, value)

	logger.Printf("INFO: result='%s' and '%s'", fqdn, value)

	return nil
}

func (d *LocalDNSProvider) CleanUp(domain, token, keyAuth string) error {

	logger.Printf("INFO: cleaning up '%s'", dns01.ToFqdn(domain))

	lookup_set("DNS01:"+domain, "")
	return nil
}

func https_init() {
	certmagic.Default.Agreed = true

	certmagic.Default.Email = *email

	certmagic.Default.CA = certmagic.LetsEncryptProductionCA

	certmagic.Default.DNSProvider = &LocalDNSProvider{}
	certmagic.Default.DisableHTTPChallenge = true
	certmagic.Default.DisableTLSALPNChallenge = true
}
