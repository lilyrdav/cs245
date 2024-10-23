defmodule Lab do
  def count(list, fun) do
   Enum.reduce(list, 0, fn(x, acc) -> if fun.(x) do acc + 1 else acc end end)
 end

 def sum(list, fun) do
    Enum.reduce(list, 0, fn(x, acc) -> if fun.(x) do acc + x else acc end end)
  end

 def max(list) do
   Enum.reduce(list, 0, fn(x, acc) -> if x > acc do x else acc end end)
 end

 def at(list, fun) do
  Enum.reduce(list, 0, fn(x, acc) -> if fun.(x) do ))
 end
end

list = [5, 6, 3]
fun = fn(x) -> x > 4 end
IO.puts(Lab.max(list, fun))
