defmodule Digraffe.Context do
  use Digraffe.Web, :model

  schema "contexts" do
    field :name, :string
    field :external_id, :string
    timestamps
  end

  @required_fields ~w(name external_id)
  @optional_fields ~w()

  def changeset(model, params \\ :empty) do
    model
    |> cast(params, @required_fields, @optional_fields)
    |> unique_constraint(:external_id)
  end

  def params_for_create(params) do
    params
    |> string_keys()
    |> ensure_external_id()
  end

  def random_string(length) do
    :crypto.strong_rand_bytes(length)
    |> Base.url_encode64()
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

  defp ensure_external_id(params) do
    if Map.has_key?(params, "external_id") do
      params
    else
      Map.put(params, "external_id", random_string(12))
    end
  end
end
