USING: kernel sequences arrays ;
IN: flatten-array

: flatten-into ( vector array -- vector )
    [
        dup array?
            [ flatten-into ]
            [ dup [ over push ] [ drop ] if ]
        if
    ] each ;

: flatten ( array -- flat )
    V{ } clone
    swap flatten-into >array ;
