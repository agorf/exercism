class Phrase
  VERSION = 1

  attr_reader :phrase

  def initialize(phrase)
    @phrase = phrase
  end

  def word_count
    phrase.scan(/[\w']+/).reduce(Hash.new(0)) do |wc, word|
      word.downcase!
      word.gsub!(/\A'|'\z/, '') # strip surrounding quotes
      wc[word] += 1
      wc
    end
  end
end
