#!/usr/bin/env bash

# Note: bash doesn't handle floats natively so we could either
# 1) use arithmetic operations and convert everything to an int
# 2) use another tool like bc or awk for those operations
# 3) treat floats as strings, in which case "03" != "3".
#
# Given that we're asked to validate the triangle and need to operate
# on the numbers, that's going to be option (2).

boolRC () { (( $? == 0 )) && echo true || echo false; }
bashTriTest () {
  shape=$1
  shift
  (( $1 <= $2 + $3 && $2 <= $1 + $3 && $3 <= $1 + $2 )) || return
  (( $1 > 0 && $2 > 0 && $3 > 0 )) || return
  case "${shape}" in
    equilateral) (( $1 == $2 && $1 == $3 ));;
    isosceles) (( $1 == $2 || $1 == $3 || $2 == $3 ));;
    scalene) (( $1 != $2 && $1 != $2 ));;
    *) echo 'Invalid'; false;;
  esac
}

triTest () {
  bcPrgm='
    define isosceles (a, b, c) { return a==b || b==c || a==c; }
    define equilateral (a, b, c) { return a==b && a==c; }
    define scalene (a, b, c) { return a!=b && b!=c && a!=c; }

    define valid (a, b, c) {
        if (a==0 || b==0 || c==0) { return 0; }
          return (a + b > c) && (b + c > a) && (a + c > b);
    }
  '
  valid=$( bc <<< "$bcPrgm valid( $2, $3, $4 )" )
  (( valid )) || return
  result=$( bc <<< "$bcPrgm $1( $2, $3, $4 )" )
  (( result ))
}


(( $# == 4 )) || exit 1
triTest "$@"
# bashTriTest "$@"
boolRC

# vim:ts=2:sw=2:expandtab
