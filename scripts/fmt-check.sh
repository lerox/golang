#!/bin/sh


GO_FMT_OUTPUT=$(gofmt -l .)
if [ -z "${GO_FMT_OUTPUT}" ]; then
  exit 0
fi

printf "Some files not following go fmt:\n%s\n" "$GO_FMT_OUTPUT"

exit 1
