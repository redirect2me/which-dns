// Copyright 2011 Miek Gieben. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net"
	"strconv"
	"strings"

	"github.com/miekg/dns"
)

func handleWhich(w dns.ResponseWriter, r *dns.Msg) {
	var (
		v4     bool
		rr     dns.RR
		str    string
		caller string
	)

	domain := r.Question[0].Name
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true
	if ip, ok := w.RemoteAddr().(*net.UDPAddr); ok {
		caller = ip.IP.String()
		str = "Port: " + strconv.Itoa(ip.Port) + " (udp) from " + caller
		v4 = ip.IP.To4() != nil
	}
	if ip, ok := w.RemoteAddr().(*net.TCPAddr); ok {
		caller = ip.IP.String()
		str = "Port: " + strconv.Itoa(ip.Port) + " (tcp) from " + caller
		v4 = ip.IP.To4() != nil
	}

	logger.Printf("INFO: DNS request for %s (%d) from %s", domain, r.Question[0].Qtype, caller)

	if v4 {
		rr = &dns.A{
			Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 0},
			A:   net.ParseIP(*ipaddress),
		}
	} else {
		rr = &dns.AAAA{
			Hdr:  dns.RR_Header{Name: domain, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 0},
			AAAA: net.ParseIP(*ipaddress),
		}
	}

	t := &dns.TXT{
		Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0},
		Txt: []string{str},
	}

	switch r.Question[0].Qtype {
	case dns.TypeAAAA, dns.TypeA:
		lookup_set(domain, caller)
		m.Answer = append(m.Answer, rr)
		m.Extra = append(m.Extra, t)
	case dns.TypeNS:
		ns := new(dns.NS)
		ns.Hdr = dns.RR_Header{Name: domain, Rrtype: dns.TypeNS, Class: dns.ClassINET, Ttl: 1}
		ns.Ns = *nshostname + "."
		m.Answer = append(m.Answer, ns)
		m.Extra = append(m.Extra, t)
	case dns.TypeSOA:
		soa, err := dns.NewRR(domain + " 0 IN SOA " + *nshostname + ". " + strings.ReplaceAll(*email, "@", ".") + ". 2019122201 86400 7200 3600000 172800")
		if err != nil {
			logger.Printf("ERROR: making soa %s", err)
		}
		m.Answer = append(m.Answer, soa)
		m.Extra = append(m.Extra, t)
	case dns.TypeTXT:
		responseTxt, ok := lookup_get("DNS01:" + domain)
		if !ok {
			logger.Printf("ERROR: no challenge answer found for %s", domain)
			lookup_debug(logger.Writer())
			responseTxt = "__ERROR_lookup_failed_"
		}
		challengeAnswer := &dns.TXT{
			Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0},
			Txt: []string{responseTxt},
		}
		m.Answer = append(m.Answer, challengeAnswer)
		m.Extra = append(m.Extra, t)
	default:
		// For numeric-> name translation, see https://github.com/miekg/dns/blob/master/types.go#L24
		// NOTE: I have seen 15 (MX) and 257 (CAA) during testing
		logger.Printf("ERROR: unknown DNS question type %d", r.Question[0].Qtype)
	}

	if *verbose {
		logger.Printf("DEBUG: DNS reply is %v", m.String())
	}

	w.WriteMsg(m)
}

func serve(proto string) {
	logger.Printf("INFO: starting DNS %s", proto)
	server := &dns.Server{Addr: "[::]:53", Net: proto, TsigSecret: nil, ReusePort: false}
	if err := server.ListenAndServe(); err != nil {
		logger.Printf("Failed to setup the dns listener for "+proto+": %s", err.Error())

	}
}

func dns_main() {

	dns.HandleFunc(".", handleWhich)
	go serve("tcp")
	go serve("udp")
}
