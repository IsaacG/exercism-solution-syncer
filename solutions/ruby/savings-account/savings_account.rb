module SavingsAccount
  private_constant

  BALANCE_TO_RATES = {
        ...0    => 3.213,
       0...1000 => 0.5,
    1000...5000 => 1.621,
    5000...     => 2.475
  }.freeze

  module_function

  def interest_rate(balance)
    BALANCE_TO_RATES.each do |range, rate|
      return rate if range.include?(balance)
    end
  end

  def annual_balance_update(balance)
    balance + balance * interest_rate(balance) / 100
  end

  def years_before_desired_balance(current_balance, desired_balance)
    (0..).each do |year|
      return year if current_balance >= desired_balance

      current_balance = annual_balance_update(current_balance)
    end
  end
end
