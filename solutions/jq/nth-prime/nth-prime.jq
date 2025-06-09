def is_prime($priors; $threshold):
  .
  | . as $candidate
  | [$priors[] | select(. <= $threshold)]
  | all($candidate % . != 0)
  ;

def next_prime:
  .
  | . as $priors
  | ((. | max) + 2 )
  | until(is_prime($priors; . | sqrt); . + 2)
  | $priors + [.]
  ;

.number as $n
| if $n < 1
  then "there is no zeroth prime" | halt_error
  else .
  end
| reduce range(1; $n - 1) as $_ ([3]; next_prime)
| .[-1]
