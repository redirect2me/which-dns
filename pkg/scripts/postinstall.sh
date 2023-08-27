#!/usr/bin/env bash

echo "INFO: running postinstall"

setcap 'cap_net_bind_service=+ep' /opt/which-dns/which-dns
systemctl daemon-reload
systemctl enable which-dns.service
systemctl start which-dns