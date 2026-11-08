#!/usr/bin/env bash

# *** PLEASE REMOVE THESE COMMENTS BEFORE SUBMITTING YOUR SOLUTION ***

main () {
echo $1 >/dev/null
  rev <<< "$1"
}

main "$@"
