#!/usr/bin/env bash
#
# make deb package for which-dns
#

set -o errexit
set -o pipefail
set -o nounset

echo "INFO: packaging started at $(date -u +%Y-%m-%dT%H:%M:%SZ)"

export VERSION=0.0.1
export PRERELEASE=$(date -u +%Y%m%dT%H%M%S)

rm -f ./dist/*.deb
nfpm pkg --packager deb --target ./dist/

echo "INFO: packaging complete at $(date -u +%Y-%m-%dT%H:%M:%SZ)"
