class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze

  class UnsupportedOperation < StandardError
  end

  def self.calculate(first_operand, second_operand, operation)
    unless operation != nil && operation != "" && ALLOWED_OPERATIONS.include?(operation)
      raise UnsupportedOperation.new
    end
    if operation == '/' && second_operand == 0
      return 'Division by zero is not allowed.'
    end
    unless first_operand.is_a?(Integer) && second_operand.is_a?(Integer)
      raise ArgumentError.new
    end
    if operation == '+'
      result = first_operand + second_operand
    elsif operation == '*'
      result = first_operand * second_operand
    elsif operation == '/'
      result = first_operand / second_operand
    end
    return "#{first_operand} #{operation} #{second_operand} = #{result}"
  end
end
