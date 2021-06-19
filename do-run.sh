#!/bin/bash
#
#

set -o errexit
set -o pipefail
set -o nounset

#
# check for required files
#
ENV_FILE="./.env"
if [ ! -f "${ENV_FILE}" ]; then
    echo "ERROR: no file '${ENV_FILE}'!"
    exit 1
fi

export $(cat ${ENV_FILE})

rm -f ~/.local/share/certmagic/locks/cert_acme_which.resolve.rs_httpsacme-v02.api.letsencrypt.orgdirectory.lock

./which-dns \
	"-disclaimer=For light, non-commerical use only!" \
	-email=${EMAIL} \
	-hostname=${HOSTNAME} \
	-ipaddress=${IPADDRESS} \
	-nshostname=${NSHOSTNAME} \
	"-tracker=${TRACKER}" \
# -local
# -verbose
