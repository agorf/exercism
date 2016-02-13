class Robot
  NAMES = {}

  def name
    return @name if @name

    @name = generate_unique_name
    NAMES[@name] = 1 # mark as taken
    @name
  end

  def reset
    NAMES.delete(@name)
    @name = nil
  end

  private

  def generate_name
    [
      ('A'..'Z').to_a.sample,
      ('A'..'Z').to_a.sample,
      rand(900) + 100
    ].join
  end

  def generate_unique_name
    name = nil

    loop do
      name = generate_name
      return name if NAMES[name].nil? # not taken
    end
  end
end
