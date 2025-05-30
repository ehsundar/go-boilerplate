#!/bin/bash
set -e

PKG_NAME=$(grep '^module ' go.mod | awk '{print $2}')
BINARY_NAME="$(basename $PKG_NAME)"
VERSION=${1:-"0.0.1"}
APP_NAME=${2:-"$BINARY_NAME"}
COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE=$(date -u +%Y-%m-%dT%H:%M:%SZ)

mkdir -p bin

GO_LDFLAGS="-X '$PKG_NAME/cmd.AppName=$APP_NAME' -X '$PKG_NAME/cmd.Version=$VERSION' -X '$PKG_NAME/cmd.Commit=$COMMIT' -X '$PKG_NAME/cmd.Date=$DATE'"

go build -ldflags="$GO_LDFLAGS" -o "bin/$BINARY_NAME" ./main.go

echo "Built bin/$BINARY_NAME version $VERSION ($COMMIT) at $DATE"
