#!/usr/bin/env bash

# Install required packages
apt-get update && apt-get install -y curl
curl -L https://github.com/gohugoio/hugo/releases/download/v0.109.0/hugo_extended_0.109.0_linux-amd64.deb -o hugo.deb
apt-get install -y ./hugo.deb
rm hugo.deb
apt-get install -y make
