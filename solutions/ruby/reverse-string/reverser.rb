class Reverser
  def self.reverse(string)
    string.grapheme_clusters.reverse.join('')
  end
end
