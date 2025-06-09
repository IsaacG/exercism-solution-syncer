"
" Convert a number to a string, the contents of which depend on the number's
" factors.
"
"   - If the number has 3 as a factor, output 'Pling'.
"   - If the number has 5 as a factor, output 'Plang'.
"   - If the number has 7 as a factor, output 'Plong'.
"   - If the number does not have 3, 5, or 7 as a factor, just pass
"     the number's digits straight through.
"

function! Raindrops(number) abort
  let out = ""
  let sounds = {3: "Pling", 5: "Plang", 7: "Plong"}
  for factor in keys(sounds)
    if a:number % factor == 0
      let out .= sounds[factor]
    endif
  endfor

  if empty(out)
    let out .= a:number
  endif

  return out
endfunction
