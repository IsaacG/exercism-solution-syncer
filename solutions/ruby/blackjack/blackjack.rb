module Blackjack
  SPLIT = 'P'.freeze
  HIT = 'H'.freeze
  STAND = 'S'.freeze
  WIN = 'W'.freeze

  private_constant

  CARD_VALUES = (
    %w[jack queen king].map { |card| [card, 10] } +
    %w[two three four five six seven eight nine ten ace].each_with_index.map { |card, index| [card, index + 2] }
  ).to_h.freeze

  module_function

  def parse_card(card)= CARD_VALUES.fetch(card, 0)

  def card_range(card1, card2)
    case [card1, card2].map { parse_card(_1) }.sum
    when 4..11 then 'low'
    when 12..16 then 'mid'
    when 17..20 then 'high'
    when 21 then 'blackjack'
    end
  end

  def first_turn(card1, card2, dealer_card)
    return SPLIT if [card1, card2].all? { |card| card == 'ace' }

    case card_range(card1, card2)
    when 'low' then HIT
    when 'mid' then parse_card(dealer_card) >= 7 ? HIT : STAND
    when 'high' then STAND
    when 'blackjack' then parse_card(dealer_card) >= 10 ? STAND : WIN
    end
  end
end
