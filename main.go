package main

import (
    "flag"
	"net"
	"strconv"
	"github.com/miekg/dns"
	"log"
)

var (
    //verbose = flag.Bool("verbose", true, "verbose logging");
    //debug = flag.Bool("debug", false, "print instead of redirect");
    //port = flag.Int("port", 80, "port to listen on");
    hostname = flag.String("hostname", "localhost", "hostname of this server");
    //action = flag.String("action", "lookup", "action [lookup|addwww|removewww]");

    //logger = log.New(os.Stdout, "R2ME: ", log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC);
)
type handler struct{}

func (this *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {

    var (
        caller string
        str string
    )

	if ip, ok := w.RemoteAddr().(*net.UDPAddr); ok {
        caller = ip.IP.String()
		str = "Address: " + ip.String() + " Port: " + strconv.Itoa(ip.Port) + " (udp)"
	}

	if ip, ok := w.RemoteAddr().(*net.TCPAddr); ok {
        caller = ip.IP.String()
		str = "Address: " + ip.String() + " Port: " + strconv.Itoa(ip.Port) + " (udp)"
    }

	msg := dns.Msg{}
	msg.SetReply(r)
    domain := msg.Question[0].Name
    lookup_set(domain, caller)
    log.Printf("INFO: resolving %s\n", domain)

	t := &dns.TXT{
		Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0},
		Txt: []string{str},
	}

	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		address := "127.0.0.1"
        msg.Answer = append(msg.Answer, &dns.A{
            Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
            A:   net.ParseIP(address),
        })
	}
    msg.Extra = append(msg.Extra, t)
	w.WriteMsg(&msg)
}

var done = make(chan bool)

func dns_main() {
	log.Printf("INFO: starting DNS\n")
	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
		done <- true
	}
}


func main() {

    flag.Parse()

	lookup_init()

	go dns_main()

	go web_main()

	<-done

	log.Printf("INFO: done\n")
}
