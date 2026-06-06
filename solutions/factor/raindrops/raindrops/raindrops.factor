USING: kernel arrays math math.functions math.parser sequences ;
IN: raindrops
CONSTANT: sounds { { 3 "Pling" } { 5 "Plang" } { 7 "Plong" } }

: convert ( n -- str )
    dup
    sounds
    [ dup first swapd divisor? [ second ] [ drop "" ] if ] with map
    "" join
    dup "" = [ drop number>string ] [ nip ] if
    ;
