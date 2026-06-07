USING: bosuns-briefing.helpers kernel arrays locals sequences vectors ;
IN: bosuns-briefing

: roster ( names -- str ) [ crew-line ] map "\n" join ;

: briefing ( names -- str )
    V{ } clone
    greeting over push
    swap roster over push
    closing over push
    "\n" join ;
