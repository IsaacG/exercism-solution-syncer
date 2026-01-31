class BankAccount
  def initialize
    @open = false
  end

  def open
    raise ArgumentError, "You can't open an already open account." if @open

    @open = true
    @balance = 0
  end

  def deposit(amnt)
    raise ArgumentError, "You can't deposit money into a closed account." unless @open
    raise ArgumentError, "You can't deposit a negative amount." if amnt.negative?

    @balance += amnt
  end

  def withdraw(amnt)
    raise ArgumentError, "You can't withdraw money into a closed account." unless @open
    raise ArgumentError, "You can't withdraw a negative amount." if amnt.negative?
    raise ArgumentError, "You can't withdraw more than you have." if amnt > @balance

    @balance -= amnt
  end

  def close
    raise ArgumentError, "You can't close an already closed account." unless @open

    @open = false
  end

  def balance
    raise ArgumentError, "You can't check the balance of a closed account." unless @open

    @balance
  end
end
