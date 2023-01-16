#!/usr/bin/env bash

# Install required packages
apt-get update && apt-get install -y make wget
wget https://github.com/gohugoio/hugo/releases/download/v0.109.0/hugo_extended_0.109.0_linux-amd64.deb
dpkg -i hugo_extended_0.109.0_linux-amd64.deb
rm -f hugo_extended_0.109.0_linux-amd64.deb
# Install markdownlint and zip tools
apt-get install zip
npm install -g markdownlint-cli
# Install golangci-lint
curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.32.2
rm -rf /usr/local/go
