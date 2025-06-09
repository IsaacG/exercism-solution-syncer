.
| if (.strand1 | length) != (.strand2 | length)
  then "strands must be of equal length" | halt_error
  else .
  end
| (.strand1 | split("")) as $s1
| (.strand2 | split("")) as $s2
| [range($s1 | length)]
| reduce .[] as $i (
    0;
    . + if $s1[$i] == $s2[$i] then 0 else 1 end
  )
| .
