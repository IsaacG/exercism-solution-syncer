USING: kernel arrays math math.parser sequences ;
IN: raindrops

: convert ( n -- str )
    dup
    { { 3 "Pling" } { 5 "Plang" } { 7 "Plong" } }
    ! [ swap [ dup first ] dip swap mod 0 = [ second ] [ drop "" ] ] with map
    [ dup first swapd mod 0 = [ second ] [ drop "" ] if ] with map
    "" join
    dup "" = [ drop number>string ] [ nip ] if
    ;
