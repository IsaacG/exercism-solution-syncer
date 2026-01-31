class Reverser
  def self.reverse(string)
    return string.grapheme_clusters.reverse.join("")
  end
end
