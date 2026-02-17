class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze
  private_constant :ALLOWED_OPERATIONS

  class UnsupportedOperation < StandardError
  end

  def self.compute(first_operand, second_operand, operation)
    case operation
    when '+'
      first_operand + second_operand
    when '*'
      first_operand * second_operand
    when '/'
      first_operand / second_operand
    end
  end
  private_class_method :compute

  def self.calculate(first_operand, second_operand, operation)
    raise UnsupportedOperation, '' unless ALLOWED_OPERATIONS.include?(operation)
    raise ArgumentError, '' unless first_operand.is_a?(Integer) && second_operand.is_a?(Integer)
    return 'Division by zero is not allowed.' if operation == '/' && second_operand.zero?

    result = compute(first_operand, second_operand, operation)
    "#{first_operand} #{operation} #{second_operand} = #{result}"
  end
end
