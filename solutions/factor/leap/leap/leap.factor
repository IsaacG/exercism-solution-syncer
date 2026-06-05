USING: kernel math ;
IN: leap

: leap-year? ( year -- ? )
    dup 4 mod 0 = swap dup 100 mod 0 = not swap 400 mod 0 = or and ;
