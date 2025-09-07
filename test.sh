#!/usr/bin/env bash
#
# show which ips match a given resolver
#

set -o errexit
set -o pipefail
set -o nounset

RESOLVER=${1:-1.1.1.1}
UUID=$(uuidgen)

dig ${UUID}.which.resolve.rs @${RESOLVER}

curl \
	--silent \
	--show-error \
	https://which.resolve.rs/debug.txt \
	| grep ${UUID} \
	| cut -f 3 -d ' '