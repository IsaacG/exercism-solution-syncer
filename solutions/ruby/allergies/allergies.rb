class Allergies
  ALLERGENS = %w[eggs peanuts shellfish strawberries tomatoes chocolate pollen cats].freeze

  private_constant :ALLERGENS

  private

  attr_writer :list

  def initialize(score)
    self.list = ALLERGENS.each_index.select { |i| ((score >> i) & 1) == 1 }.map { |i| ALLERGENS[i] }
  end

  public

  attr_reader :list

  def allergic_to?(allergen)
    list.include?(allergen)
  end
end
