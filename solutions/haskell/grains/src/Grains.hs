module Grains (square, total) where

import Control.Monad

square :: Integer -> Maybe Integer
square n = do 
  guard (1 <= n && n <= 64)
  return (2 ^ (n - 1))

total :: Integer
total = 2 ^ 64 - 1
