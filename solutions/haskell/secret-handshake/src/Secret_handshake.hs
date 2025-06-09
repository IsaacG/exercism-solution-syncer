module SecretHandshake (handshake) where

import Data.Bits

construct :: Int -> [String]
construct n =
  map
    snd
    $ filter
      (testBit n . fst)
      $ zip [0..] actions
  where actions = ["wink", "double blink", "close your eyes", "jump"]

handshake :: Int -> [String]
handshake n
  | n .&. 16 == 0 = construct n
  | otherwise     = reverse $ construct n

-- vim:ts=2:sw=2:expandtab
