defmodule Digraffe.Context do
  use Digraffe.Web, :model
  import Digraffe.Util

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
end
