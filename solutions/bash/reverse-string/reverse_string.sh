#!/usr/bin/env bash

main () {
  echo $1 >/dev/null
  rev <<< "$1"
}

main "$@"
