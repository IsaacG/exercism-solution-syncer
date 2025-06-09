def is_prime($priors; $threshold):
  def under:
    if $priors[.] > $threshold
      then empty
      else $priors[.], (. + 1 | under)
    end
  ;

  . as $candidate |
  [0 | under] |
  all($candidate % . != 0)
;

def is_prime_alt($priors; $threshold):
  def _is_prime:
    . as [$idx, $candidate] |
    if $priors[$idx] > $threshold
     then true
      elif $candidate % $priors[$idx] == 0
      then false
      else [$idx + 1, $candidate] | _is_prime
    end
  ;

  [0, .] | _is_prime
;

def next_prime:
  . |
  . as $priors |
  ((. | max) + 2 ) |
  until(is_prime($priors; sqrt); . + 2) |
  $priors + [.]
;

def nth_prime:
  .number as $n |
  if $n < 1
    then "there is no zeroth prime" | halt_error
    elif $n == 1
    then 2
    else
      reduce range(1; $n - 1) as $_ ([3]; next_prime) |
      .[-1]
  end
;

nth_prime
