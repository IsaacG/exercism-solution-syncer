module Blackjack
  SPLIT = 'P'.freeze
  HIT = 'H'.freeze
  STAND = 'S'.freeze
  WIN = 'W'.freeze

  TEN_FACE_CARDS = %w[jack queen king].freeze
  SOME_CARDS = %w[two three four five six seven eight nine ten ace].freeze
  private_constant :TEN_FACE_CARDS, :SOME_CARDS

  def self.parse_card(card)
    if TEN_FACE_CARDS.include?(card)
      10
    elsif SOME_CARDS.include?(card)
      SOME_CARDS.index(card) + 2
    else
      0
    end
  end

  def self.card_range(card1, card2)
    case [card1, card2].map { parse_card(_1) }.sum
    when 4..11 then 'low'
    when 12..16 then 'mid'
    when 17..20 then 'high'
    when 21 then 'blackjack'
    else raise 'error'
    end
  end

  def self.first_turn(card1, card2, dealer_card)
    return SPLIT if card1 == 'ace' && card2 == 'ace'

    case card_range(card1, card2)
    when 'low' then HIT
    when 'mid' then parse_card(dealer_card) >= 7 ? HIT : STAND
    when 'high' then STAND
    when 'blackjack' then parse_card(dealer_card) >= 10 ? STAND : WIN
    end
  end
end
