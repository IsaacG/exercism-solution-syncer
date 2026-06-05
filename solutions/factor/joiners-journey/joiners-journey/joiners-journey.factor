USING: kernel math ;
IN: joiners-journey

: with-kerf ( length -- length+kerf )
    [ 0.02 * ] keep + >integer ;

: kerf-and-finish ( length -- kerf finish )
    dup [ 2 * 100 / ] [ 5 * 100 / ] bi* ;

: cut-card ( length -- length kerf finish )
    dup kerf-and-finish ;

: per-piece ( bolt-length pieces -- per-piece )
    [ with-kerf ] dip / ;

: compare-bolts ( length-a length-b -- kerf-a kerf-b )
    [ 2 * 100 / ] bi@ ;
