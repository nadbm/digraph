defmodule Digraffe.Collection do
  use Digraffe.Web, :model

  schema "collections" do
    field :title, :string
    belongs_to :owner, Digraffe.Settings
    timestamps
  end

  @required_fields ~w(title owner_id)
  @optional_fields ~w()

  def changeset(model, params \\ :empty) do
    model
    |> cast(params, @required_fields, @optional_fields)
  end
end
