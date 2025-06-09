.legacy
| to_entries
| [
  .[]
  | .key as $k
  | .value[]
  | {
      key: (. | ascii_downcase),
      value: ($k | tonumber)
    }
  ]
| sort
| from_entries
