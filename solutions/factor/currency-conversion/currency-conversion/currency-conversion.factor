USING: kernel math math.order ;
IN: currency-conversion

: exchange-money ( budget exchange-rate -- exchanged ) / ;

: get-change ( budget exchanging-value -- change ) - ;

: value-of-bills ( denomination number-of-bills -- value ) * ;

: number-of-bills ( amount denomination -- bills ) /i ;

: leftover-of-bills ( amount denomination -- leftover ) mod ;

: exchangeable-value ( denomination budget spread exchange-rate -- value )
    ! effective-rate = (spread / 100 + 1) * exchange-rate
    swap 100 / 1 + *
    ! stack: denomination budget effective-rate
    exchange-money
    ! stack: denomination amount
    dupd swap number-of-bills
    ! stack: denomination bills
    value-of-bills
    ;

: safe-change ( budget exchanging-value -- change )
    get-change 0 max ;

: cap-spend ( budget price -- spend )
    min ;
