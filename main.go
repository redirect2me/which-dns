package main

import (
    "fmt"
	"net"
    "net/http"
    "strconv"
    //"sync"
	"log"
	"github.com/miekg/dns"
)

var domainsToAddresses map[string]string = map[string]string{
	"google.com.": "1.2.3.4",
	"jameshfisher.com.": "104.198.14.52",
}

type handler struct{}
func (this *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		log.Printf("INFO: resolving %s\n", domain);
		address, ok := "127.0.0.1", true //domainsToAddresses[domain]
		if ok {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{ Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60 },
				A: net.ParseIP(address),
			})
		}
	}
	w.WriteMsg(&msg)
}

var done = make(chan bool)


func dns_main() {
	log.Printf("INFO: starting DNS\n");
	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
        done <- true
	}
}

func web_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func web_main() {
    http.HandleFunc("/", web_handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

    //wg.Add(1);
    go dns_main();

    //wg.Add(1);
    go web_main();

    <-done

	log.Printf("INFO: done\n");
}