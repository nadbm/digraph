defmodule Digraffe.Link do
  use Digraffe.Web, :model

  alias Digraffe.Util

  schema "links" do
    field :title,       :string
    field :url,         :string
    field :external_id, :string
    timestamps
  end

  @required_fields ~w(title external_id url)
  @optional_fields ~w()

  def changeset(model, params \\ :empty) do
    model
    |> cast(params, @required_fields, @optional_fields)
    |> unique_constraint(:external_id)
  end

  def params_for_create(%{url: _url} = params) do
    params
    |> Util.string_keys()
    |> params_for_create()
  end

  def params_for_create(%{"url" => url} = params) do
    Util.params_for_create(params, fn ->
      {_status, string} = Util.normalize_url(url)
      :crypto.hash(:sha, string)
    end)
  end

  def params_for_create(params) do
    params
  end
end
