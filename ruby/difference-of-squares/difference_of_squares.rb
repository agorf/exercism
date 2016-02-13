class Squares
  VERSION = 2

  attr_reader :n

  def initialize(n)
    @n = n
  end

  def square_of_sum
    (1..n).reduce(:+).to_i ** 2
  end

  def sum_of_squares
    (1..n).map {|i| i ** 2 }.reduce(:+).to_i
  end

  def difference
    square_of_sum - sum_of_squares
  end
end
