package main

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mholt/certmagic"
)

type apiResult struct {
	Success    bool   `json:"success"`
	Timestamp  string `json:"timestamp"`
	Message    string `json:"message,omitempty"`
	Input      string `json:"input,omitempty"`
	Output     string `json:"output,omitempty"`
	Source     string `json:"source"`
	Disclaimer string `json:"disclaimer,omitempty"`
}

func getRemoteIP(req *http.Request) string {
	if *proxy {
		remoteIPs := req.Header.Get("X-Forwarded-For")
		comma := strings.Index(remoteIPs, ",")
		if comma == -1 {
			return remoteIPs
		}
		return remoteIPs[0:comma]
	}
	remoteIP, _, _ := net.SplitHostPort(req.RemoteAddr)
	return remoteIP
}

func api_handler(w http.ResponseWriter, r *http.Request) {
	result := apiResult{}
	result.Timestamp = time.Now().UTC().Format(time.RFC3339)
	result.Disclaimer = *disclaimer

	callback := r.FormValue("callback")
	if callback == "" {
		result.Success = false
		result.Message = "Missing callback parameter"
		logger.Printf("WARNING: no callback parameter")
		write_with_callback(w, r, result)
		return
	}

	hostname := r.Host
	if strings.Contains(hostname, ":") {
		var err error
		hostname, _, err = net.SplitHostPort(r.Host)
		if err != nil {
			result.Input = r.Host
			result.Success = false
			result.Message = "Unable to split host/port"
			logger.Printf("WARNING: lookup failed %s (%s)", result.Message, r.Host)
			write_with_callback(w, r, result)
			return
		}
	}

	if hostname == "" {
		result.Success = false
		result.Message = "URL has no host"
		logger.Printf("WARNING: lookup failed %s (%s/%s)", result.Message, r.Host, r.URL.String())
		write_with_callback(w, r, result)
		return
	}

	hostname += "."
	result.Input = hostname

	result.Output, result.Success = lookup_get(hostname)
	if !result.Success {
		result.Message = "hostname not found in cache"
		logger.Printf("WARNING: lookup failed %s (%s)", result.Message, hostname)
	}
	result.Source = getRemoteIP(r)

	write_with_callback(w, r, result)
}

func debug_handler(w http.ResponseWriter, r *http.Request) {
	lookup_debug(w)
}

func favicon_ico_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/vnd.microsoft.icon")
	w.Write(favicon_ico)
}

func favicon_svg_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf8")
	w.Write([]byte(`<svg width="64" height="64" version="1.1" viewBox="0 0 64 64" xmlns="http://www.w3.org/2000/svg"><g transform="matrix(.125 0 0 .125 0 -3.3336)"><path d="m400.3 198.95-170.67-106.66c-3.156-1.977-7.109-2.156-10.438-0.492-3.312 1.672-5.531 4.953-5.843 8.664-2.938 35.726 5.265 64.373 24.39 85.139 12.859 13.96 30.546 24.155 52.577 30.304 16.516 4.617 35.438 6.961 56.233 6.961 27.937 0 49.296-4.227 50.202-4.406 4.281-0.867 7.609-4.25 8.391-8.547 0.781-4.306-1.141-8.642-4.844-10.962z" fill="#434a54"/><path d="m495.57 360.23c-7.781-4.094-18.718-7.141-33.405-9.297-27.953-4.093-67.796-4.655-106.67-4.655-14.655 0-29.904 0.094-46.045 0.188-17.25 0.094-35.108 0.188-53.452 0.188-18.359 0-36.218-0.094-53.467-0.188-16.14-0.094-31.375-0.188-46.046-0.188-38.874 0-78.701 0.562-106.67 4.655-14.688 2.156-25.609 5.203-33.406 9.297-13.562 7.141-16.406 16.86-16.406 23.765 0 11.688 7.953 23.375 23.671 34.797 13.437 9.766 32.546 19.234 55.233 27.391 51.124 18.374 114.01 28.468 177.09 28.468 63.062 0 125.95-10.094 177.07-28.468 22.688-8.156 41.797-17.625 55.233-27.391 15.718-11.422 23.687-23.109 23.687-34.797 1e-3 -6.905-2.858-16.64-16.42-23.765z" fill="#434a54"/><path d="m362.3 386.63-64.764-248.47c-0.688-2.641-2.359-4.914-4.672-6.359l-63.232-39.515c-2.938-1.836-6.578-2.133-9.766-0.797-3.203 1.344-5.546 4.148-6.296 7.523l-63.999 287.99c-0.172 0.766-0.25 1.547-0.25 2.328 0 3.688 1.781 13.047 18.328 20.406 17.297 7.688 47.015 11.594 88.342 11.594 41.312 0 71.03-3.905 88.326-11.594 16.546-7.359 18.327-16.719 18.327-20.406-1e-3 -0.921-0.11-1.812-0.344-2.702z" fill="#656d78"/><g fill="#ffce54"><path d="m355.43 360.24c-2.641 2.312-6.233 4.656-11.108 6.828-17.296 7.688-47.015 11.578-88.326 11.578-41.327 0-71.045-3.891-88.342-11.578-5.422-2.406-9.25-5.031-11.953-7.609l-6.125 27.546c-0.172 0.766-0.25 1.547-0.25 2.328 0 3.688 1.781 13.047 18.328 20.406 17.297 7.688 47.015 11.594 88.342 11.594 41.312 0 71.03-3.905 88.326-11.594 16.546-7.359 18.327-16.719 18.327-20.406 0-0.922-0.109-1.812-0.344-2.703z"/><rect x="235.32" y="378.65" width="42.67" height="42.28"/></g><path d="m189 409.74c-16.562-7.359-18.344-16.719-18.344-20.406 0-0.781 0.094-1.562 0.25-2.328l63.999-287.99c0.219-0.984 0.594-1.906 1.062-2.773l-6.328-3.953c-2.938-1.836-6.578-2.133-9.766-0.797-3.203 1.344-5.546 4.148-6.296 7.523l-63.999 287.99c-0.172 0.766-0.25 1.547-0.25 2.328 0 3.688 1.781 13.047 18.328 20.406 17.297 7.688 47.015 11.594 88.342 11.594 3.641 0 7.188-0.031 10.656-0.094-35.796-0.625-61.889-4.499-77.654-11.5z" enable-background="new" fill="#fff" opacity=".1"/><path d="m277.99 431.6h-42.67c-5.891 0-10.672-4.766-10.672-10.672v-42.28c0-5.875 4.781-10.656 10.672-10.656h42.67c5.891 0 10.656 4.781 10.656 10.656v42.28c0 2.844-1.125 5.547-3.125 7.547s-4.703 3.125-7.531 3.125zm-31.999-21.328h21.328v-20.938h-21.328z" fill="#f6bb42"/></g></svg>`))
}

func spinner_svg_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml; charset=utf8")
	w.Write([]byte(`<svg id="svg-zone-spinner" viewBox="0 0 50 50" width="100%" height="100%" xmlns="http://www.w3.org/2000/svg">
	<g stroke="black" stroke-width="2" stroke-linecap="round">
		<path opacity=".3" d="M29 18.071797 L33, 11.143594"/>
		<path opacity=".3"  d="M31.928203 21 L38.856406 17" />
		<path opacity=".3" d="M25 33 L25 41" />
		<path opacity=".3" d="M31.928203 29 L38.856406 33, "/>
		<path opacity=".3"  d="M29 31.928203 L33, 38.856406"/>
		<path opacity=".3" d="M33 25 L41 25" />
		<path opacity=".3"  d="M21 31.928203 L17, 38.856406"/>
		<path opacity=".44"  d="M18.071797 29 L11.143594 33" />
		<path opacity=".65" d="M17 25 L9 25" />
		<path opacity=".86"  d="M18.071797 21 L11.143594 17, "/>
		<path opacity=".93"  d="M21 18.071797 L17, 11.143594"/>
		<path opacity="1.0" d="M25 17 L25 9" />
		<animateTransform attributeType="xml"
						  attributeName="transform"
						  type="rotate"
						  values="0 25 25; 30 25 25; 60 25 25; 90 25 25; 120 25 25; 150 25 25; 180 25 25; 210 25 25; 240 25 25; 270 25 25; 300 25 25; 330 25 25"
						  calcMode="discrete"
						  dur="1.0s"
						  repeatCount="indefinite"/>
	</g>
</svg>`))
}

func base_robots_txt_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.Write([]byte(`#
# robots.txt for which-dns
#
# Just a single page, but feel free to index it!
#
User-Agent: *
Allow: /
`))
}

func leaf_robots_txt_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf8")
	w.Write([]byte(`#
# robots.txt for which-dns leaf nodes (*.` + *hostname + `)
#
# go to https://` + *hostname + `/ for more information
#
User-Agent: *
Disallow: /
`))
}

func root_handler(w http.ResponseWriter, r *http.Request) {

	reqhostname := r.Host
	if strings.Contains(reqhostname, ":") {
		reqhostname, _, _ = net.SplitHostPort(r.Host)
	}
	if reqhostname != *hostname {
		http.Redirect(w, r, "https://"+*hostname+"/", http.StatusPermanentRedirect)
		return
	}

	if r.URL.Path[1:] != "" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	trackerEl := ""

	if *tracker != "" {
		trackerEl = "\n\t\t" + `<img alt="" src="` + *tracker + `" />`
	}
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
	<head>
	<meta charset="utf-8">
		<title>Which DNS</title>
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<link rel="icon" type="image/svg+xml" href="/favicon.svg" />
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/kognise/water.css@latest/dist/light.min.css" />
	</head>
	<body>
		<h1>
			<img alt="Which DNS logo" src="favicon.svg" style="height:2.2em;vertical-align:middle;" />
			Which DNS
		</h1>
		<p>
			Determine which DNS servers you are actually using.
		</p>
		<p>
			Your DNS Server:
			<span id="ipresult"><img src="/images/spinner.svg" style="height:16pt;" /></span>
			<span id="dnsresult" style="display:none;"><img src="/images/spinner.svg" style="height:16pt;" /></span>
		</p>
		<p>
			<a href="https://github.com/redirect2me/which-dns">How this works</a>, including source code!
		</p>
		<p>
			More troubleshooting tools are available at <a href="https://resolve.rs/tools.html">resolve.rs</a>.
		</p>
		<script>
			function updateResolver(data) {
				console.log('DNS resolver data', data);
				var el = document.getElementById("ipresult");
				if (data.success) {
					el.innerText = data.output + (data.output == data.source ? " (self-hosted)" : "");
					var asnel = document.getElementById("dnsresult");
					asnel.style.display = 'inline';
					var head = document.getElementsByTagName('head')[0];
					var script = document.createElement('script');
					script.type = 'text/javascript';
					script.src = 'https://resolve.rs/dns/reverse-lookup.json?callback=updateReverseDns&ip=' + data.output;
					head.appendChild(script);
					console.log(asnel);
				} else {
					el.innerText = data.message;
				}
			}

			function updateReverseDns(data) {
				console.log('Reverse DNS data', data);
				var el = document.getElementById("dnsresult");
				if (data.success && data.results && data.results.length > 0) {
					el.innerText = data.asn + ": " + data.results[0];
				} else {
					el.innerText = data.asn;
				}
			}
		</script>
		<script src="https://` + uuid.New().String() + `.which.resolve.rs/api.json?callback=updateResolver"></script>` + trackerEl + `
	</body>
</html>`))

}

func web_main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", root_handler)
	mux.HandleFunc("/favicon.ico", favicon_ico_handler)
	mux.HandleFunc("/favicon.svg", favicon_svg_handler)
	mux.HandleFunc("/images/spinner.svg", spinner_svg_handler)
	mux.HandleFunc(*hostname+"/robots.txt", base_robots_txt_handler)
	mux.HandleFunc("/robots.txt", leaf_robots_txt_handler)
	mux.HandleFunc("/status.json", status_handler)
	mux.HandleFunc("/debug.txt", debug_handler)
	mux.HandleFunc("/api.json", api_handler)

	lm := LoggingMiddleware(*logger)
	//mux = lm(mux)

	if *local {
		logger.Printf("WARNING: running locally w/o https")
		logger.Fatal(http.ListenAndServe(":4000", lm(mux)))
	} else {
		logger.Printf("INFO: initializing https")
		https_init()
		logger.Printf("INFO: starting web server on %s", *hostname)
		certmagic.HTTPS([]string{"*." + *hostname, *hostname}, mux)
	}
}
