#!/usr/bin/env sh

SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"

docker run --rm -it -v "$SCRIPTPATH:/go/src/github.com/ns-consulting/fizz-buzz-rest-server" -w='/go/src/github.com/ns-consulting/fizz-buzz-rest-server' golang:1.10.3-alpine3.8 go "$@"
