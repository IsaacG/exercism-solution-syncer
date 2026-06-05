USING: kernel combinators unicode math ;
IN: character-study

SYMBOLS: less equal greater
    big small no-size
    alpha numeric space newline unknown ;

: compare-chars ( c1 c2 -- symbol )
    {
        { [ 2dup < ] [ 2drop less ] }
        { [ 2dup > ] [ 2drop greater ] }
        [ 2drop equal ]
    } cond ;

: size-of-char ( c -- symbol )
    {
        { [ dup LETTER? ] [ drop big ] }
        { [ dup letter? ] [ drop small ] }
        [ drop no-size ]
    } cond ;

: change-size-of-char ( c desired -- c' )
    {
        { [ dup big = ] [ drop ch>upper ] }
        { [ dup small = ] [ drop ch>lower ] }
        [ drop ]
    } cond ;

: type-of-char ( c -- symbol )
    {
        { [ dup Letter? ] [ drop alpha ] }
        { [ dup digit? ] [ drop numeric ] }
        { [ dup 32 = ] [ drop space ] }
        { [ dup CHAR: \n = ] [ drop newline ] }
        [ drop unknown ]
    } cond ;
