#!/usr/bin/env bash
#
# build which-dns
#

set -o errexit
set -o pipefail
set -o nounset

echo "INFO: build started at $(date -u +%Y-%m-%dT%H:%M:%SZ)"

SCRIPT_HOME="$( cd "$( dirname "$0" )" && pwd )"
APP_HOME=$(realpath "$SCRIPT_HOME/..")

LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION=${1:-local@$(date -u +%Y-%m-%dT%H:%M:%SZ)}
GITHUB_SHA=${GITHUB_SHA:-$(git rev-parse --short HEAD)}

echo "PWD=$(pwd)"
echo "APP_HOME=${APP_HOME}"
echo "SCRIPT_HOME=${SCRIPT_HOME}"
echo "GITHUB_SHA=${GITHUB_SHA:-(not set)}"

echo "INFO: creating directory"
mkdir -p "${SCRIPT_HOME}/dist"

echo "INFO: building"
go build \
    -a \
    -trimpath \
    -ldflags "-s -w -extldflags '-static' -X main.COMMIT=$GITHUB_SHA -X main.LASTMOD=$LASTMOD -X main.VERSION=$VERSION" \
    -installsuffix cgo \
    -tags netgo \
    -o "${APP_HOME}/dist/which-dns" \
    "${APP_HOME}/src"

echo "INFO: build complete at $(date -u +%Y-%m-%dT%H:%M:%SZ)"