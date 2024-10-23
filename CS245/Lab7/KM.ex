defmodule KM do
    def main() do
        IO.inspect(System.argv())
        number = Float.parse(hd(System.argv))
        unit = hd(tl(System.argv))
        km = number * 1.609
        mi = number / 1.609
        end
        
    end
end
