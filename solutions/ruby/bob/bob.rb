# frozen_string_literal: true

module Bob
  def self.hey(sentence)
    sentence.strip!
    yelling = sentence == sentence.upcase && sentence != sentence.downcase
    return "Calm down, I know what I'm doing!" if yelling && sentence.end_with?('?')
    return 'Sure.' if sentence.end_with?('?')
    return 'Whoa, chill out!' if yelling
    return 'Fine. Be that way!' if sentence.empty?

    'Whatever.'
  end
end
