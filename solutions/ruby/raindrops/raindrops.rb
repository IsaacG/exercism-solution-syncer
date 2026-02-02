module Raindrops
  SOUNDS = { 3 => 'Pling', 5 => 'Plang', 7 => 'Plong' }.freeze

  def self.convert(number)
    sound = SOUNDS.select { |factor, _| (number % factor).zero? }.values.join('')
    sound.empty? ? number.to_s : sound
  end
end
