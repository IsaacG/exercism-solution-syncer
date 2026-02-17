class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze
  private_constant :ALLOWED_OPERATIONS

  class UnsupportedOperation < StandardError
  end

  def self.calculate(first_operand, second_operand, operation)
    raise UnsupportedOperation, '' unless ALLOWED_OPERATIONS.include?(operation) && first_operand.respond_to?(operation)
    raise ArgumentError, '' unless first_operand.is_a?(Integer) && second_operand.is_a?(Integer)
    return 'Division by zero is not allowed.' if operation == '/' && second_operand.zero?

    result = first_operand.send(operation, second_operand)
    "#{first_operand} #{operation} #{second_operand} = #{result}"
  end
end
