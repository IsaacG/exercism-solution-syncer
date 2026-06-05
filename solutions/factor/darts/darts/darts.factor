USING: kernel math combinators ;
IN: darts

: score ( x y -- n )
    sq swap sq +
    {
        { [ dup 100 > ] [ drop  0 ]  }
        { [ dup 25 >  ] [ drop  1 ]  }
        { [ dup 1 >   ] [ drop  5 ]  }
        [ drop 10 ]
    } cond ;
    
