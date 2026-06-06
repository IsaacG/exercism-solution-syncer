USING: kernel math math.statistics sequences ;
IN: librarians-ledger

: protected-balance ( opening requests -- balance )
    swap [ + dup neg? [ drop 0 ] [ ] if ] reduce ;

: running-balance ( transactions -- balances )
    cum-sum ;

: least-balance-so-far ( transactions -- worsts )
    running-balance cum-min ;

: halve-until ( principal target -- balances )
    swap [ 2dup < ] [ 2 /i dup ] produce 2nip ;
