module ProteinTranslation(proteins) where

protein :: String -> Maybe String
protein "AUG" = Just "Methionine"
protein "UUC" = Just "Phenylalanine"
protein "UUU" = Just "Phenylalanine"
protein "UUG" = Just "Leucine"
protein "UUA" = Just "Leucine"
protein "UCG" = Just "Serine"
protein "UCA" = Just "Serine"
protein "UCC" = Just "Serine"
protein "UCU" = Just "Serine"
protein "UAC" = Just "Tyrosine"
protein "UAU" = Just "Tyrosine"
protein "UGC" = Just "Cysteine"
protein "UGU" = Just "Cysteine"
protein "UGG" = Just "Tryptophan"
protein "UAG" = Just "Stop"
protein "UAA" = Just "Stop"
protein "UGA" = Just "Stop"
protein _ = Nothing

proteins :: String -> Maybe [String]
proteins "" = Just []
proteins s = case protein first of
  Nothing -> Nothing
  Just "Stop" -> Just []
  Just amino -> case proteins remaining of
    Nothing -> Nothing
    Just rest -> Just $ amino : rest
  where
    (first, remaining) = splitAt 3 s
