class Sieve
  def initialize(n)
    @n = n
  end

  attr_reader :n

  def primes
    nums = (2..n).to_a

    nums.each_with_index do |prime, i|
      next if prime.nil?

      nums[i + 1..-1].each.with_index(i + 1) do |num, j|
        next if num.nil?

        nums[j] = nil if num % prime == 0
      end
    end

    nums.compact
  end
end
