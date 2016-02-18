defmodule Acronym do
  @doc """
  Generate an acronym from a string. 
  "This is a string" => "TIAS"
  """
  @spec abbreviate(string) :: String.t()
  def abbreviate(string) do
    Regex.scan(~r/[A-Z]|(?<= )[a-z]/, string)
      |> List.flatten
      |> List.to_string
      |> String.upcase
  end
end
