USING: kernel arrays sequences vectors ;
IN: mosaic-mischief

: fresh-mosaic ( n -- row )
    f <array> ;

: place-tile ( row position colour -- )
    spin set-nth ;

: chip-tile ( row position -- )
    f place-tile ;

: recolour-tile ( row position quot -- )
    swapd change-nth ; inline

: snapshot-mosaic ( row -- copy )
    clone ;

: stash-tile ( hoard tile -- )
    swap push ;

: return-tile ( hoard -- tile )
    pop ;
