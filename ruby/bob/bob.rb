class Bob
  def hey(phrase)
    case phrase.gsub(/[^\w?]/, '') # clean phrase
    when /\A[A-Z0-9]*[A-Z][A-Z0-9]*\??\z/
      'Whoa, chill out!'
    when /\?\z/
      'Sure.'
    when ''
      'Fine. Be that way!'
    else
      'Whatever.'
    end
  end
end
