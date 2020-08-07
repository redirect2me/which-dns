#!/bin/bash

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

COMMIT=$(git rev-parse --short HEAD)-local
LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)

echo "INFO: compiling"
GOOS=linux GOARCH=amd64 go build \
    -ldflags "-X main.COMMIT=${COMMIT} -X main.LASTMOD=${LASTMOD}" \
    .

echo "INFO: copying to server"
scp -i ~/.ssh/do do-run.sh .env root@${IPADDRESS}:
scp -i ~/.ssh/do which-dns root@${IPADDRESS}:which-dns.new

echo "INFO: done!"
