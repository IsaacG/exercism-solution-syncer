#!/usr/bin/env bash

(( $# == 4 )) || exit 1
declare -ri cap1=$1 cap2=$2 goal=$3

check_complete () {
  local b1=$1 b2=$2 moves=$3
  (( b1 == goal || b2 == goal )) || return
  if (( b1 > 0 && b1 < cap1 )); then
    goalBucket=one other=$b2
  else
    goalBucket=two other=$b1
  fi
  printf 'moves: %d, goalBucket: %s, otherBucket: %d\n' $moves $goalBucket $other
  exit
}

fill_pour_empty_append () {
  local b1=$1 b2=$2 moves=$3 swap=$4 c no_fill_pour=0

  # Rather than doing b1,b2 use full|empty bucket volume|capacity.
  (( b1 )) && (( fv=b1, fc=cap1, ev=0, ec=cap2 )) || (( fv=b2, fc=cap2, ev=0, ec=cap1 ))

  # Pour from one bucket into the other prior to filling and pouring.
  if (( swap )); then
    if (( fv > ec )); then # Full bucket volume does not fit neatly in empty.
      (( ev=ec, fv-=ev, moves++ ))
      (( no_fill_pour++ ))
    else
      (( c=b1, b1=b2, b2=c, c=fc, fc=ec, ec=c, moves++ ))
    fi
  fi

  (( space = fc-fv ))
  if (( no_fill_pour )); then
    :
  elif (( space == 0 )); then # We got a full and empty bucket. No filling. Pour full into empty.
    (( fc > ec )) && (( ev=ec, fv-=ec, moves++ )) || (( ev=fv, fv-=ev, moves++ ))
  elif (( space == ec )); then # we are just filling up the partially filled bucket
    (( fv+=ec, ev=0, moves++ ))
  else
    if (( space > ec )); then  # fill, pour, two moves
      (( fv+=ec, ev=0 ))
    else
      (( fv=fc, ev=ec-space ))
    fi
    (( moves += 2 ))
  fi

  # Convert back from full/empty to b1/b2
  (( b1 )) && (( b1=fv, b2=ev, 1 )) || (( b2=fv, b1=ev ))

  check_complete $b1 $b2 $moves
  (( b1 == cap1 )) && (( b1=0, moves++ ))
  (( b2 == cap2 )) && (( b2=0, moves++ ))
  echo $b1 $b2 $moves
}

buckets () {
  local -i b1 b2 moves=1
  local -a to_explore explored
  local start=$4

  # Initial conditions
  [[ $start = 'one' ]] && (( b2=0, b1=cap1 )) || (( b1=0, b2=cap2 ))
  # Pre-check if we are done before we start.
  check_complete $b1 $b2 $moves
  check_complete $cap1 $b2 2
  check_complete $b1 $cap2 2

  # We did not yet explore anything.
  for (( i = 0; i <= cap1 || i <= cap2; i++ )); do explored[i]=0; done

  # Situations to start our exploring: the initial setup.
  to_explore=( "$b1 $b2 $moves" )

  while (( "${#to_explore[@]}" )); do
    read b1 b2 moves <<< "${to_explore[0]}"
    to_explore=( "${to_explore[@]:1}" )

    (( explored[b1 ? b1 : b2] )) && continue || explored[b1 ? b1 : b2]=1
    [[ $start = 'two' ]] && (( b1 == cap1 && b2 == 0 )) && continue
    # Explore with and without an initial transfer
    next=$( fill_pour_empty_append $b1 $b2 $moves 0 ) && to_explore+=( "$next" )
    [[ $next = moves* ]] && echo "$next" && exit
    next=$( fill_pour_empty_append $b1 $b2 $moves 1 ) && to_explore+=( "$next" )
    [[ $next = moves* ]] && echo "$next" && exit
    printf -v exp '[%s] ' "${to_explore[@]}"
  done
  echo "invalid goal"
  exit 1
}

buckets "$@"

# vim:ts=2:sw=2:expandtab
