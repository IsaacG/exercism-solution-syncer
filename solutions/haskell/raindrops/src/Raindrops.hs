module Raindrops (convert) where

addSound :: Int -> Int -> String -> String
addSound num factor sound
  | mod num factor == 0 = sound
  | otherwise = ""

convert :: Int -> String
convert n = if s == "" then show n else s
  where s = (addSound n 3 "Pling") ++ (addSound n 5 "Plang") ++ (addSound n 7 "Plong") 
