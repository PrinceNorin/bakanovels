#!/bin/sh

find . -type f \( -name "*.go" -o -name "*.json" -o -name "*.yml" \) | entr -r ./run.sh
