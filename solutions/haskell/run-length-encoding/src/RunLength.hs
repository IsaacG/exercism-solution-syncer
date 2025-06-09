module RunLength (decode, encode) where

import Data.Char

toInt :: String -> Int
toInt s = read s :: Int

toChar :: Int -> String
toChar s = show s :: String

decode :: String -> String
decode (x:xs)
  | n /= "" = replicate (toInt n) (head r) ++ (decode . tail $ r)
  | otherwise = x:decode xs
  where
    (n, r) = span isDigit $ x:xs
decode [] = ""

encode :: String -> String
encode s
  | length s <= 1 = s
  | length h == 1 = c:encode t
  | otherwise = toChar (length h) ++ [c] ++ encode t
    where
      (h, t) = span (== head s) s
      c = head s

-- vim:ts=2:sw=2:expandtab
