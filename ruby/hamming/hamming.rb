module Hamming
  VERSION = 1

  def self.compute(strand1, strand2)
    raise ArgumentError if strand1.size != strand2.size

    distance = 0

    strand1.each_char.with_index do |c, i|
      distance += 1 if strand2[i] != c
    end

    distance
  end
end
