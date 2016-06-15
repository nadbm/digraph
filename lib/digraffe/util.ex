defmodule Digraffe.Util do

  @id_length 12

  def params_for_create(params, seed \\ &random_string/0) do
    string_id = id_string(seed.())
    params
    |> string_keys()
    |> ensure_external_id(string_id)
  end

  def id_string(binary) do
    binary
    |> Base.url_encode64()
    |> simple_chars()
    |> binary_part(0, @id_length)
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

  @defaults %{
    path:   "/",
    scheme: "http",
  }

  def normalize_url(string) do
    case parse_uri(string) do
      {:ok,    uri} -> {:ok, normalize_uri(uri)}
      {:error, _}   -> {:error, string}
    end
  end

  defp normalize_uri(uri) do
    Map.merge(@defaults, uri, fn
      _k,  v1, nil -> v1
      _k, _v1, v2  -> v2
    end)
    |> Map.to_list()
    |> downcase_values()
    |> URI.to_string()
  end

  defp parse_uri(string) do
    uri = URI.parse(string)
    case uri do
      %URI{host: nil} ->
        {:error, uri}
      %URI{host: host} ->
        if String.contains?(host, ".") do
          {:ok, uri}
        else
          {:error, uri}
        end
      _ ->
        {:error, uri}
    end
  end

  defp downcase_values(parts) do
    Enum.reduce(parts, %URI{}, fn
      {_k, nil}, acc                -> acc
      {:__struct__, _v}, acc        -> acc
      {k, v}, acc when is_binary(v) -> Map.put(acc, k, String.downcase(v))
      {k, v}, acc                   -> Map.put(acc, k, v)
    end)
  end

  defp random_string() do
    :crypto.strong_rand_bytes(@id_length + 10)
  end

  defp ensure_external_id(params, external_id) do
    if Map.has_key?(params, "external_id") do
      params
    else
      Map.put(params, "external_id", external_id)
    end
  end
end
