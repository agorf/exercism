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

  def self.primes_until(n)
    primes = (2..n).to_a

    primes.each_with_index do |prime, i|
      next if prime.nil?

      (i + prime...primes.size).step(prime).each do |j|
        num = primes[j]

        next if num.nil?

        if num % prime == 0
          primes[j] = nil
        end
      end
    end

    primes.compact
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
end
