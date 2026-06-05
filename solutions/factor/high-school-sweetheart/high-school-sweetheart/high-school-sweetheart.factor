USING: kernel splitting sequences strings ascii ;
IN: high-school-sweetheart

: cleanupname ( name -- cleaned )
    "-" " " replace [ blank? ] trim ;

: firstletter ( name -- letter )
    cleanupname first 1string ;

: initial ( name -- initial )
    firstletter >upper "." append ;

: couple ( name1 name2 -- formatted )
    swap initial swap initial "  +  " glue " " dup surround "\u{2764}" dup surround ;
