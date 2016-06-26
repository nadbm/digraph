defmodule Digraffe.Collection do
  use Digraffe.Web, :model

  @primary_key {:id, :binary_id, autogenerate: true}

  schema "collections" do
    field :title, :string
    timestamps
  end

  @required_fields ~w(title)
  @optional_fields ~w()

  def changeset(model, params \\ :empty) do
    model
    |> cast(params, @required_fields, @optional_fields)
  end
end
