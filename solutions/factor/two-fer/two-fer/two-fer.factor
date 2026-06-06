USING: kernel formatting ;
IN: two-fer

! There are no variadic functions in Factor, due to the nature of the stack.

: 2-for-1 ( name/f -- str )
    [ "you" ] unless*
    "One for %s, one for me." sprintf ;
