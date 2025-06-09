def colors:
  ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"]
  ;

[
  .colors[0:2][]
  | . as $c
  | colors
  | index($c)
]
| reduce .[] as $i (0; . * 10 + $i)
