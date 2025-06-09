#!/usr/bin/env bash

encode () {
  local n=$1 in=$2 d=1 r=0
  local -a rails
  for (( i = 0; i < ${#in}; i++ )); do
    rails[$r]+=${in:i:1}
    (( r += d ))
    (( (r == 0) || (r == n - 1) )) && (( d *= -1 ))
  done
  printf '%s' "${rails[@]}"; printf \\n
}


(( $# == 3 )) || { echo "Invalid usage"; exit 1; }
(( $2 > 0 )) || { echo "Invalid rail count"; exit 1; }

case "$1" in
  -e) encode "$2" "$3";;
  -d) decode "$2" "$3";;
  *) exit 1;;
esac

# vim:ts=2:sw=2:expandtab
