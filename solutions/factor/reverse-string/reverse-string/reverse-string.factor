USING: kernel sequences strings ;
IN: reverse-string

: reverse-string ( str -- str )
    dup "" = [ ] [ unclip [ reverse-string ] dip 1string append ] if ;
