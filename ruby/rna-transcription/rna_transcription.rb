module Complement
  VERSION = 3

  def self.of_dna(strand)
    raise ArgumentError unless strand =~ /\A[GCTA]+\z/
    strand.tr('GCTA', 'CGAU')
  end
end
