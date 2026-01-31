module Blackjack
  SPLIT = 'P'.freeze
  HIT = 'H'.freeze
  STAND = 'S'.freeze
  WIN = 'W'.freeze

  def self.parse_card(card)
    case card
    when 'two' then 2
    when 'three' then 3
    when 'four' then 4
    when 'five' then 5
    when 'six' then 6
    when 'seven' then 7
    when 'eight' then 8
    when 'nine' then 9
    when 'ten', 'jack', 'queen', 'king' then 10
    when 'ace' then 11
    else 0
    end
  end

  def self.card_range(card1, card2)
    case [card1, card2].map { |card| parse_card(card) }.sum
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
