module ListOps
  def self.arrays(array)= array.length

  def self.reverser(array)= array.reverse

  def self.concatter(one, two)= one + two

  def self.filterer(array)= array.filter { |i| yield(i) }

  def self.mapper(array)= array.map { |i| yield(i) }

  def self.sum_reducer(array)= array.sum

  def self.factorial_reducer(array)
    out = 1
    array.each { |i| out *= i }
    out
  end
end
