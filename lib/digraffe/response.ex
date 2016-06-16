defmodule Digraffe.Http.Response do

  def title(response) do
    case find(response.headers, "Content-Type") do
      {"Content-Type", "text/html"} ->
        case Floki.find(response.body, "title") do
          [{"title", _attrs, values}] ->
            Enum.at(values, 0)
          _ -> nil
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
