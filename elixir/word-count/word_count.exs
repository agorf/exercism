defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t) :: map()
  def count(sentence) do
    words = Regex.split(~r/[^[:alnum:]-]+/u, String.downcase(sentence), [trim: true])
    count(%{}, words)
  end

  defp count(freq, [head|tail]) do
    new_count = Map.get(freq, head, 0) + 1
    new_freq = Map.put(freq, head, new_count)
    count(new_freq, tail)
  end

  defp count(freq, []) do
    freq
  end
end
