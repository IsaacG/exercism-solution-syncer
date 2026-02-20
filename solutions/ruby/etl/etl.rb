module ETL
  def self.transform(data)
    [].union(*data.map do |value, keys|
      keys.map { |key| [key.downcase, value] }
    end).to_h
  end
end
