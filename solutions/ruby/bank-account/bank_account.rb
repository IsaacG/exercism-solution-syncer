class BankAccount
  def initialize
    @open = false
  end

  def open
    if @open
      raise ArgumentError.new("You can't open an already open account.")
    end
    @open = true
    @balance = 0
  end

  def deposit(amnt)
    unless @open
      raise ArgumentError.new("You can't deposit money into a closed account.")
    end
    if amnt < 0
      raise ArgumentError.new("You can't deposit a negative amount.")
    end

    @balance += amnt
  end

  def withdraw(amnt)
    unless @open
      raise ArgumentError.new("You can't withdraw money into a closed account.")
    end
    if amnt < 0
      raise ArgumentError.new("You can't withdraw a negative amount.")
    end
    if amnt > @balance
      raise ArgumentError.new("You can't withdraw more than you have.")
    end

    @balance -= amnt
  end

  def close
    unless @open
      raise ArgumentError.new("You can't close an already closed account.")
    end

    @open = false
  end

  def balance
    unless @open
      raise ArgumentError.new("You can't check the balance of a closed account.")
    end

    @balance
  end
end

