#!/usr/bin/env bash

echo "INFO: starting which-dns"



CONFIG_FILE="${XDG_CONFIG_HOME:-/etc/}which-dns.conf"

if [ -f "${CONFIG_FILE}" ]; then
    echo "INFO: loading config file '${CONFIG_FILE}'"
    source "${CONFIG_FILE}"
fi

#HOSTNAME=$(hostname --fqdn)
#IPADDRESS=$(dig +short txt ch whoami.cloudflare @1.0.0.1)
#LOG=/var/log/	$XDG_STATE_HOME/log/

/opt/which-dns/which-dns \
    -bind "159.203.79.70" \
    -disclaimer "Pray for me!" \
    -email "admin@example.com" \
    -hostname "tyche" \
    -ipaddress 127.0.0.1 \
    -local true \
    -nshostname "which-pkg.resolve.rs" \
    -proxy false \
    -verbose true

