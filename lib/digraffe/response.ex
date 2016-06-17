defmodule Digraffe.Http.Response do

  def title(response) do
    case find(response.headers, "Content-Type") do
      {"Content-Type", all_types} ->
        types = String.split(all_types, ~r{;\s*}) |> Enum.map(fn s -> String.downcase(s) end)
        if Enum.member?(types, "text/html") do
          case Floki.find(response.body, "title") do
            [{"title", _attrs, values}] ->
              Enum.at(values, 0)
            _ -> nil
          end
        else
          nil
        end
      _ -> nil
    end
  end

  def find(headers, header_name) do
    Enum.find(headers, nil, fn
      {name, _value} -> header_name == name
    end)
  end
end
