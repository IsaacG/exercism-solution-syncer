module Grains (square, total) where

square :: Integer -> Maybe Integer
square n 
  | n > 64 || n < 1 = Nothing
  | otherwise = Just (2 ^ (n - 1))

total :: Integer
total = 2 ^ 64 - 1
