class BaseConverter
  def self.validate(input_base, digits, output_base)
    raise ArgumentError if input_base < 2
    raise ArgumentError if output_base < 2
    raise ArgumentError if digits.any? { |digit| digit >= input_base || digit.negative? }
  end

  def self.combine_input(input_base, digits)
    input = 0
    digits.each { |digit| input = input * input_base + digit }
    input
  end

  def self.generate_output(input, output_base)
    out = []
    until input.zero?
      out.append(input % output_base)
      input /= output_base
    end
    out.append(0) if out.empty?
    out.reverse
  end

  def self.convert(input_base, digits, output_base)
    validate(input_base, digits, output_base)
    generate_output(combine_input(input_base, digits), output_base)
  end
end
