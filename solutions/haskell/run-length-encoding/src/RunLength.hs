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
    n = takeWhile isDigit $ x:xs
    r = dropWhile isDigit $ x:xs
decode [] = ""

encode :: String -> String
encode s
  | length s <= 1 = s
  | n == 1 = h:encode t
  | otherwise = toChar n ++ [h] ++ encode t
    where
      h = head s
      t = dropWhile (== head s) s
      n = length . takeWhile  (== head s) $ s

-- vim:ts=2:sw=2:expandtab
