module Grains
  def self.square(n)
    2**(n - 1)
  end

  def self.total
    #(1..64).map {|n| square(n) }.reduce(:+)
    64.times.reduce(0) {|total, _| 2 * total + 1 }
  end
end
