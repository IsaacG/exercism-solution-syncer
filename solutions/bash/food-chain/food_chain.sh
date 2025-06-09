#!/usr/bin/env bash

declare -a creatures=(fly spider bird cat dog goat cow horse)
declare -A reaction=(
  [spider]='It wriggled and jiggled and tickled inside her.'
  [bird]='How absurd to swallow a bird!'
  [cat]='Imagine that, to swallow a cat!'
  [dog]='What a hog, to swallow a dog!'
  [goat]='Just opened her throat and swallowed a goat!'
  [cow]="I don't know how she swallowed a cow!"
  [horse]="She's dead, of course!"
)
declare -A reason=(
  [cow]='She swallowed the cow to catch the goat.'
  [goat]='She swallowed the goat to catch the dog.'
  [dog]='She swallowed the dog to catch the cat.'
  [cat]='She swallowed the cat to catch the bird.'
  [bird]='She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.'
  [spider]='She swallowed the spider to catch the fly.'
  [fly]="I don't know why she swallowed the fly. Perhaps she'll die."
)

print_verse () {
  local -i i
  local animal=${creatures[$1-1]}

  printf 'I know an old lady who swallowed a %s.\n' "$animal"
  [[ ${reaction[$animal]} ]] && printf '%s\n' "${reaction[$animal]}"
  [[ $animal = 'horse' ]] && exit
  for (( i = $1 - 1; i >= 0; i-- )); do
    [[ ${reason[${creatures[$i]}]} ]] && printf '%s\n' "${reason[${creatures[$i]}]}"
  done
}

song () {
  local -i i

  if ! (( $# == 1 || $# == 2 )); then
    echo '1 or 2 arguments expected'
    exit 1
  fi

  start=$1
  end=${2:-$start}

  if (( start > end )); then
    echo 'Start must be less than or equal to End'
    exit 1
  fi

  for (( i = start; i < end; i++ )); do
    print_verse "$i"
    printf '\n'
  done
  print_verse "$end"
}

song "$@"


# vim:ts=2:sw=2:expandtab
