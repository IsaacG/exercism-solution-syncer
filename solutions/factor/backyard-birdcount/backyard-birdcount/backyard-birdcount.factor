USING: kernel sequences math ;
IN: backyard-birdcount

: today ( days -- count/f )
    [ f ] [ first ] if-empty ; 

: increment-day-count ( days -- new-days )
    [ { 1 } ] [ unclip 1 + prefix ] if-empty ;

: has-day-without-birds? ( days -- ? )
    [ f ] [ unclip 0 = [ has-day-without-birds? ] dip or ] if-empty ;

: total ( days -- sum )
    [ 0 ] [ unclip [ total ] dip + ] if-empty ;

: busy-days ( days -- count )
    [ 0 ] [ unclip [ busy-days ] dip 5 < [ 0 ] [ 1 ] if + ] if-empty ;
