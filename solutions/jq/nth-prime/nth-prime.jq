def is_prime:
  .
  | .candidate as $candidate
  | ($candidate | sqrt | ceil) as $threshold
  | .priors
  | [.[] | select(. <= $threshold)]
  | all($candidate % . != 0)
  ;

def next_prime:
  .
  | .candidate |= . + 2
  | if is_prime
    then (.priors += [.candidate])
    else next_prime
    end
  ;

.number as $n
| if $n < 1
  then "there is no zeroth prime" | halt_error
  else .
  end
| reduce range(1; $n) as $_ ({candidate: 1, priors: [2]}; next_prime)
| .priors[$n - 1]
