USING: kernel sequences splitting ascii ;
IN: log-levels

: message ( log-line -- message )
    ": " split1 nip  [ blank? ] trim ;

: log-level ( log-line -- level )
    "[]" split second >lower ;

: reformat ( log-line -- formatted )
    dup message swap log-level "(" ")" surround " " glue ;
