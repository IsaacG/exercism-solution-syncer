USING: kernel math ;
IN: eliuds-eggs

: egg-count ( n -- count )
    dup 0 =                        ! if n == 0
    [ ]                            ! return 0
    [                              ! else return 0|1 + egg-count(n >> 1)
        [ 0 bit? [ 1 ] [ 0 ] if ]
        [ -1 shift egg-count ]
        bi +
    ]
    if
    ;
