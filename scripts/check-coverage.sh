#!/usr/bin/env bash
set -euo pipefail

go test ./... -coverprofile=coverage.out

total=$(go tool cover -func=coverage.out | awk '/total:/ {print substr($3, 1, length($3)-1)}')
threshold=90

if [[ -z "$total" ]]; then
  echo "coverage: unable to read total" >&2
  exit 1
fi

result=$(awk -v total="$total" -v threshold="$threshold" 'BEGIN {if (total+0 >= threshold) print "ok"; else print "fail"}')
if [[ "$result" != "ok" ]]; then
  echo "coverage ${total}% is below ${threshold}%" >&2
  exit 1
fi

echo "coverage ${total}%"
