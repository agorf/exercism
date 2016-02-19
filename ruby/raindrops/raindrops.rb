module Raindrops
  VERSION = 1

  SOUNDS = {
    3 => 'Pling',
    5 => 'Plang',
    7 => 'Plong',
  }

  def self.convert(n)
    str = (prime_factors_of(n) & SOUNDS.keys).map {|pf| SOUNDS[pf] }.join
    str.empty? ? n.to_s : str
  end

  def self.prime_factors_of(n)
    primes    = primes_until(n)
    factors   = []
    remainder = n
    i         = 0

    while remainder > 1 do
      prime = primes[i]

      if remainder % prime == 0
        factors << prime
        remainder /= prime
      else
        i += 1
      end
    end

    factors
  end

  # Sieve of Eratosthenes
  def self.primes_until(n)
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
