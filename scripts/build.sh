#!/usr/bin/env sh

if [ -z "$1" ]; then
  os=`go env GOOS`
else
  os=$1
fi

if [ -z "$2" ]; then
  arch=`go env GOARCH`
else
  arch=$2
fi

echo "building for $os/$arch"
export GOOS="$os"
export GOARCH="$arch"
go build -o "./.build/$GOOS-$GOARCH/" "./cmd/ketch-event-forwarder/main.go"
