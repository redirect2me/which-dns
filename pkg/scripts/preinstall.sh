#!/usr/bin/env bash

echo "INFO: running preinstall"
id which-dns || useradd -s /sbin/nologin -M which-dns

#LATER: create log directory
