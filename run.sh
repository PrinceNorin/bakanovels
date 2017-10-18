#!/bin/sh

pkill bakanovels && echo "Sent kill"
rm -f ./bakanovels && echo "Remove old binary"

echo "Build & Run..."
go build && ./bakanovels
