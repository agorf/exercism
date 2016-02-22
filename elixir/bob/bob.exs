defmodule Bob do
  def hey(input) do
    clean_input = Regex.replace(~r/[^\w?]/u, input, "")

    cond do
      String.ends_with?(clean_input, "?") ->
        "Sure."
      String.match?(clean_input, ~r/\A[^a-z]+\z/) && !String.match?(clean_input, ~r/\A\d+\z/) ->
        "Whoa, chill out!"
      clean_input == "" ->
        "Fine. Be that way!"
      true ->
        "Whatever."
    end
  end
end
