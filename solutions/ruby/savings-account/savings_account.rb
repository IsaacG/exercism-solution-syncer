module SavingsAccount
  RATES = { 0 => 3.213, 1000 => 0.5, 5000 => 1.621 }.freeze
  MAX_RATE = 2.475

  def self.interest_rate(balance)
    (RATES.to_a.find { |threshold, _| balance < threshold } || [0, MAX_RATE])[1]
  end

  def self.annual_balance_update(balance)
    balance + balance * interest_rate(balance) / 100
  end

  def self.years_before_desired_balance(current_balance, desired_balance)
    years = 0
    while current_balance < desired_balance
      years += 1
      current_balance = annual_balance_update(current_balance)
    end
    years
  end
end
