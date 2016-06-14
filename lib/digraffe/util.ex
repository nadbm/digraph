defmodule Digraffe.Util do

  def params_for_create(params) do
    params
    |> string_keys()
    |> ensure_external_id()
  end

  def id_string(length) do
    :crypto.strong_rand_bytes(length + 10)
    |> Base.url_encode64()
    |> simple_chars()
    |> binary_part(0, length)
  end

  def string_keys(map) do
    for {key, val} <- map, into: %{} do
      cond do
        is_atom(key) -> {to_string(key), val}
        true         -> {key, val}
      end
    end
  end

  def simple_chars(string) do
    string
    |> String.replace(~r/[-_=]/, "")
  end

  defp ensure_external_id(params) do
    if Map.has_key?(params, "external_id") do
      params
    else
      Map.put(params, "external_id", id_string(12))
    end
  end
end
