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

./which-dns -hostname=${HOSTNAME} -ipaddress=${IPADDRESS}
# -local