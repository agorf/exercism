class SumOfMultiples
  def self.to(n)
    new(3, 5).to(n)
  end

  def initialize(*nums)
    @nums = nums
  end

  attr_reader :nums

  def to(n)
    multiples = []

    nums.each do |num|
      i = 1

      while (multiple = num * i) < n
        multiples << multiple
        i += 1
      end
    end

    multiples.uniq.reduce(:+) || 0
  end
end
