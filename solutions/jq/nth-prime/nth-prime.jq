def is_prime($priors):
  def _is_prime:
    .
    | . as [$idx, $threshold, $candidate]
    | if $priors[$idx] > $threshold
      then true
      elif $candidate % $priors[$idx] == 0
      then false
      else [$idx + 1, $threshold, $candidate] | _is_prime
      end
    ;

  [0, sqrt, .] | _is_prime
  ;

def next_prime:
  .
  | . as $priors
  | ((. | max) + 2 )
  | until(is_prime($priors); . + 2)
  | $priors + [.]
  ;

.number as $n
| if $n < 1
  then "there is no zeroth prime" | halt_error
  elif $n == 1
  then 2
  else
    .
    | reduce range(1; $n - 1) as $_ ([3]; next_prime)
    | .[-1]
  end
