#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

COMMIT=$(git rev-parse --short HEAD)-local
LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)

go run \
    -ldflags "-X main.COMMIT=${COMMIT} -X main.LASTMOD=${LASTMOD}" \
    main.go lookup.go web_server.go favicon_ico.go jsonp.go status.go \
    -local
