USING: kernel ascii strings ranges sequences ;
IN: pangram

: pangram? ( sentence -- ? )
    >lower
    CHAR: a CHAR: z [a..b] [ swap member? ] with all?
    ;
