class SumOfMultiples
  def self.to(n)
    new(3, 5).to(n)
  end

  def initialize(*nums)
    @nums = nums
  end

  def multiple?(n)
    @nums.any? {|num| n % num == 0 }
  end
  private :multiple?

  def to(n)
    (0...n).select {|num| multiple?(num) }.reduce(:+)
  end
end
