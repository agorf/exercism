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
    ints = (3..n).step(2).to_a
    p = 3

    loop do
      (p**2..n).step(2*p).each {|num| ints[num - 2] = nil }
      break unless p = ints.find {|num| num && num > p }
    end

    ints.compact!
    ints.unshift(2) if n > 1
    ints
  end
end
