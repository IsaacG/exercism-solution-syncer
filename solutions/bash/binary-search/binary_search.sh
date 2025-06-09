#!/usr/bin/env bash

(( $# )) || exit 1

# Recursive solution.
# Note this can be solved with a while/until-loop
# which tends to be faster and more efficient.
b_find () {
  local -i want=$1 start=$2 end=$3
  shift 3
  local -a array=( "$@" )
  
  local -i midpoint=$(( (start + end) / 2 ))
  local -i mpv=${array[midpoint]} # midpoint val

  if (( mpv == want )); then # the value is at the midpoint
    echo "$midpoint"
  elif (( start >= end )); then # we exhausted the array
    echo -1
  elif (( mpv < want )); then # value is after the midpoint
    b_find "$want" "$((midpoint + 1))" "$end" "${array[@]}"
  elif (( mpv > want )); then # value is before the midpoint
    b_find "$want" "$start" "$((midpoint - 1))" "${array[@]}"
  else
    echo "How did you get here?"
    exit 1
  fi
}

search () {
  local want=$1
  shift
  local start=0 end="$(( $# - 1 ))"

  b_find "$want" "$start" "$end" "$@"
}

search "$@"

# vim:ts=2:sw=2:expandtab
