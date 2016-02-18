defmodule RunLengthEncoder do 
  @doc """
  Generates a string where consecutive elements are represented as a data value and count.
  "HORSE" => "1H1O1R1S1E"
  For this example, assume all input are strings, that are all uppercase letters.
  It should also be able to reconstruct the data into its original form. 
  "1H1O1R1S1E" => "HORSE" 
  """
  @spec encode(string) :: String.t
  def encode(string) do
    Regex.scan(~r/([A-Z])\1*/, string)
      |> Enum.map(fn([whole, c]) -> [String.length(whole), c] |> Enum.join end)
      |> Enum.join
  end

  @spec decode(string) :: String.t
  def decode(string) do
    Regex.scan(~r/(\d+)([A-Z])/, string)
      |> Enum.map(fn([_, n, str]) -> String.duplicate(str, String.to_integer(n)) end)
      |> Enum.join
  end
end
