class Array
  def accumulate(&blk)
    ret = []
    each {|el| ret << yield(el) }
    ret
  end
end
