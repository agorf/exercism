class Binary
  VERSION = 1

  def initialize(num)
    raise ArgumentError unless num =~ /\A[01]+\z/

    @num = num
    @size = num.size
  end

  attr_reader :num, :size

  def to_decimal
    num.each_char.with_index.reduce(0) do |sum, (char, i)|
      sum + char.to_i * 2**(size - i - 1)
    end
  end
end
