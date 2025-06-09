module Raindrops (convert) where


convert :: Int -> String
convert n = if null result then show n else result
  where
    getSound :: Int -> Int -> String -> String
    getSound num factor sound = if num `mod` factor == 0 then sound else ""

    result = getSound n 3 "Pling" ++ getSound n 5 "Plang" ++ getSound n 7 "Plong"