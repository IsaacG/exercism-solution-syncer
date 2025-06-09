#!/usr/bin/env bash

fail () {
    echo "Error: invalid input"
    exit 1
}

grains () {
    (( $# == 1 )) || fail
    if [[ "$1" = "total" ]]; then
        result="$(( (2 ** 64) - 1 ))"
    else
        (( $1 >= 1 && $1 <= 64 )) || fail
        result="$(( 2 ** ("$1" - 1) ))"
    fi
    printf "%u\n" "$result"
}

grains "$@"