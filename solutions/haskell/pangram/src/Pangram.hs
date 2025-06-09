module Pangram (isPangram) where

import qualified Data.Char as C
import qualified Data.Set as S


foundAll :: S.Set Char -> S.Set Char -> String -> Bool
foundAll s t (c:cs) = t `S.isSubsetOf` s || foundAll (S.insert c s) t cs
foundAll s t []     = t `S.isSubsetOf` s


isPangram :: String -> Bool
isPangram text = foundAll S.empty asciiLower lower_input
  where
    lower_input = map C.toLower text
    asciiLower = S.fromList ['a'..'z']

-- vim:ts=2:sw=2:expandtab
