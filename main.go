package main

import (
	"flag"
	"log"
	"os"
)

var (
	//port = flag.Int("port", 80, "port to listen on");

	disclaimer = flag.String("disclaimer", "", "disclaimer text including in API responses")
	email      = flag.String("email", "", "email for Let's Encrypt account")
	hostname   = flag.String("hostname", "localhost", "hostname of this server")
	ipaddress  = flag.String("ipaddress", "127.0.0.1", "public ip address of this server")
	local      = flag.Bool("local", false, "local development (=no https, high ports)")
	nshostname = flag.String("nshostname", "localhost-ns", "hostname used as the authoritative DNS server (NS record)")
	proxy      = flag.Bool("proxy", false, "behind a trusted proxy (for recording source IP address)")
	tracker    = flag.String("tracker", "", "URL of tracking pixel")
	verbose    = flag.Bool("verbose", true, "verbose logging")

	logger = log.New(os.Stdout, "WHICH-DNS: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
)

var done = make(chan bool)

func main() {

	flag.Parse()
	if *verbose {
		logger.Printf("DEBUG: disclaimer  = %s", *disclaimer)
		logger.Printf("DEBUG: email       = %s", *email)
		logger.Printf("DEBUG: hostname    = %s", *hostname)
		logger.Printf("DEBUG: ipaddress   = %s", *ipaddress)
		logger.Printf("DEBUG: nshostname  = %s", *nshostname)
		logger.Printf("DEBUG: tracker     = %s", *tracker)
		logger.Printf("DEBUG: verbose     = %v", *verbose)
	}

	lookup_init()

	go dns_main()

	go web_main()

	<-done

	logger.Printf("INFO: done\n")
}
