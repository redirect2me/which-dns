#!/usr/bin/env bash
#
# run locally
#
# NOTE: this may need to be run with 'sudo'
#

set -o errexit
set -o pipefail
set -o nounset

COMMIT=$(git rev-parse --short HEAD)-local
LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)

go run \
    -ldflags "-X main.COMMIT=${COMMIT} -X main.LASTMOD=${LASTMOD}" \
    . \
    -local
