module ProteinTranslation(proteins) where

protein :: String -> Maybe String
protein s
  | s == "AUG" = Just "Methionine"
  | s == "UUC" = Just "Phenylalanine"
  | s == "UUU" = Just "Phenylalanine"
  | s == "UUG" = Just "Leucine"
  | s == "UUA" = Just "Leucine"
  | s == "UCG" = Just "Serine"
  | s == "UCA" = Just "Serine"
  | s == "UCC" = Just "Serine"
  | s == "UCU" = Just "Serine"
  | s == "UAC" = Just "Tyrosine"
  | s == "UAU" = Just "Tyrosine"
  | s == "UGC" = Just "Cysteine"
  | s == "UGU" = Just "Cysteine"
  | s == "UGG" = Just "Tryptophan"
  | otherwise = Nothing

proteins :: String -> Maybe [String]
proteins s = Just $ translate s

translate :: String -> [String]
translate "" = []
translate s = case p of
  Nothing ->  []
  Just q ->  q : (translate $ drop 3 s)
  where
    p = protein $ take 3 s
