class Affine
  NUM = ('0'..'9').to_a
  ALNUM = ('a'..'z').to_a + NUM

  private_constant :NUM, :ALNUM

  private

  attr_accessor :key_a, :key_b, :mmi

  def initialize(key_a, key_b)
    raise ArgumentError, 'a must be coprime with 26' if key_a.gcd(26) != 1

    self.key_a = key_a
    self.key_b = key_b
    self.mmi = (1..26).find { |n| (key_a * n) % 26 == 1 }
  end

  def encode_char(char)
    NUM.include?(char) ? char : ALNUM[(key_a * ALNUM.index(char) + key_b) % 26]
  end

  def decode_char(char)
    NUM.include?(char) ? char : ALNUM[mmi * (ALNUM.index(char) - key_b) % 26]
  end

  public

  def encode(plaintext)
    plaintext
      .downcase
      .chars
      .select { |c| ALNUM.include?(c) }
      .map { |c| encode_char(c) }
      .each_slice(5)
      .map(&:join)
      .join(' ')
  end

  def decode(ciphertext)
    ciphertext
      .downcase
      .chars
      .select { |c| ALNUM.include?(c) }
      .map { |c| decode_char(c) }
      .join
  end
end
