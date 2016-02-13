module Complement
  VERSION = 3

  SWAPS = {
    'G' => 'C',
    'C' => 'G',
    'T' => 'A',
    'A' => 'U',
  }

  def self.of_dna(strand)
    strand.chars.map do |c|
      raise ArgumentError unless compl = SWAPS[c]
      compl
    end.join
  end
end
