class InvalidCodonError < StandardError
end

module Translation
  CODONS = {
    'AUG' => 'Methionine',
    'UUU' => 'Phenylalanine',
    'UUC' => 'Phenylalanine',
    'UUA' => 'Leucine',
    'UUG' => 'Leucine',
    'UCU' => 'Serine',
    'UCC' => 'Serine',
    'UCA' => 'Serine',
    'UCG' => 'Serine',
    'UAU' => 'Tyrosine',
    'UAC' => 'Tyrosine',
    'UGU' => 'Cysteine',
    'UGC' => 'Cysteine',
    'UGG' => 'Tryptophan',
    'UAA' => :STOP,
    'UAG' => :STOP,
    'UGA' => :STOP
  }.freeze
  private_constant :CODONS

  def self.of_rna(strand)
    strand
      .chars
      .each_slice(3)
      .map(&:join)
      .take_while { CODONS.fetch(_1, :INVALID) != :STOP }
      .map do |codon|
        raise InvalidCodonError unless CODONS.include?(codon)

        CODONS[codon]
      end
  end
end
