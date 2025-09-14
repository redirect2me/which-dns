#!/usr/bin/env bash
#
# script to run on the server to complete the deployment
#

set -o errexit
set -o pipefail
set -o nounset

if [ ! -f ./which-dns.new ]; then
    echo "ERROR: no file 'which-dns.new'!"
    exit 1
fi

# stop the current service
echo "INFO: stopping current service"
killall which-dns || true

# move the new binary into place
echo "INFO: moving new binary into place"
mv ./which-dns.new ./which-dns

echo "INFO: setting execute permissions"
chmod +x ./which-dns

echo "INFO: stopping screen"
screen -S which-dns -X quit || true

echo "INFO: starting new service"
screen -d -m -S which-dns ./do-run.sh
