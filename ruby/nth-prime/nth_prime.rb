class Prime
  def self.nth(n)
    raise ArgumentError if n == 0

    return 2 if n == 1

    i = 2

    (3..Float::INFINITY).step(2).each do |num|
      num = num.to_i

      if is_prime?(num)
        return num if i == n
        i += 1
      end
    end
  end

  # Source: https://en.wikipedia.org/wiki/Primality_test#Pseudocode
  def self.is_prime?(n)
    return false if n <= 1
    return true  if n <= 3
    return false if n % 2 == 0 || n % 3 == 0

    i = 5

    while i**2 <= n
      return false if n % i == 0 || n % (i + 2) == 0
      i += 6
    end

    true
  end
end
