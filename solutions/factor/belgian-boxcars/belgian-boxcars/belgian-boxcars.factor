USING: kernel grouping sequences splitting splitting.monotonic ;
IN: belgian-boxcars

: couple ( cars n -- trains )
    group ;

: peek-couplings ( cars -- pairs )
    2 clump ;

: split-at-junctions ( cars junctions -- legs )
    swap [ dupd swap member? ] split-when nip ;

: coalesce-cargo ( cars -- runs )
    [ = ] monotonic-split ;
