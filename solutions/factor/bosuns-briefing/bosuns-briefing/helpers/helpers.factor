USING: kernel sequences ;
IN: bosuns-briefing.helpers

: greeting ( -- str ) "All hands, attention!" ;

: crew-line ( name -- str ) "  - " prepend ;

: closing ( -- str ) "Carry on." ;
