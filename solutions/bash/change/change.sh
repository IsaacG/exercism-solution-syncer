#!/usr/bin/env bash

target=$1; shift
coins=( $@ )

# Catch a negative value.
if (( target < 0 )); then
  echo "target can't be negative"
  exit 1
fi

# Compute the optimal coins/count for each value up until target.
# Apply dynamic programming. Try using each coin N and see if count(target-N)+1 is the min count.
# count(target) = min( count(target-N)+1 for N in coins )
# using(target) stores the coin N for which we got to the min.
count=(0) using=(0)
for (( i = 1; i <= target; i++ )); do
  min=-1
  for n in "${coins[@]}"; do
    (( n > i )) && continue # ignore coins larger than target
    maybe=$(( ${count[i-n]} + 1 ))
    (( maybe )) || continue
    if (( min == -1 )) || (( maybe < min )); then
      min=$maybe
      count[i]=$maybe
      using[i]=$n
    fi
  done
  (( min == -1 )) && (( count[i] = -1, using[i] = -1 ))
done

# Check if we were able to solve it.
if (( ${count[target]} == -1 )); then
  echo "can't make target with given coins"
  exit 1
fi

# Walk back using the used[target] list, accumulating the coins we used.
used=()
while (( target > 0 )); do
   used=( "${using[target]}" "${used[@]}" )
   (( target -= "${using[target]}" ))
done

# Sort the used list.
sorted=()
for n in "${coins[@]}"; do
  for u in "${used[@]}"; do
    (( n == u )) && sorted+=( $n )
  done
done
echo "${sorted[@]}"

# vim:ts=2:sw=2:expandtab
