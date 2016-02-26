defmodule ListOps do
  # Please don't use any external modules (especially List) in your
  # implementation. The point of this exercise is to create these basic functions
  # yourself.
  #
  # Note that `++` is a function from an external module (Kernel, which is
  # automatically imported) and so shouldn't be used either.

  @spec count(list) :: non_neg_integer
  def count(list), do: do_count(list, 0)
  defp do_count([], total), do: total
  defp do_count([_ | tail], total) do
    do_count(tail, total + 1)
  end

  @spec reverse(list) :: list
  def reverse(list), do: do_reverse(list, [])
  defp do_reverse([], reversed), do: reversed
  defp do_reverse([head|tail], reversed) do
    do_reverse(tail, [head | reversed])
  end

  @spec map(list, (any -> any)) :: list
  def map(list, func), do: do_map(reverse(list), func, [])
  defp do_map([], _func, result), do: result
  defp do_map([head|tail], func, result) do
    do_map(tail, func, [func.(head) | result])
  end

  @spec filter(list, (any -> as_boolean(term))) :: list
  def filter(list, func), do: do_filter(reverse(list), func, [])
  defp do_filter([], _func, result), do: result
  defp do_filter([head|tail], func, result) do
    do_filter(
      tail,
      func,
      (if func.(head), do: [head | result], else: result)
    )
  end

  @type acc :: any
  @spec reduce(list, acc, ((any, acc) -> acc)) :: acc
  def reduce([], acc, _func), do: acc
  def reduce([head|tail], acc, func) do
    reduce(tail, func.(head, acc), func)
  end

  @spec append(list, list) :: list
  def append(a, b), do: do_append(reverse(a), b)
  defp do_append([], b), do: b
  defp do_append([head|tail], b) do
    do_append(tail, [head | b])
  end

  @spec concat([[any]]) :: [any]
  def concat(ll), do: do_concat(reverse(ll), [])
  defp do_concat([], result), do: result
  defp do_concat([head|tail], result) do
    do_concat(tail, append(head, result))
  end
end
