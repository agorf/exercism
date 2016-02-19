class Array
  def accumulate
    ret = []
    each {|el| ret << yield(el) }
    ret
  end
end
