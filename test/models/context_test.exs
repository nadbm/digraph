defmodule Digraffe.ContextTest do
  use Digraffe.ModelCase
  import Digraffe.Factory
  alias Digraffe.Context

  @valid_attrs %{name: "Home", external_id: "value"}
  @invalid_attrs %{}

  test "changeset with valid attributes" do
    context = create(:context)
    params = Context.params_for_create(@valid_attrs)
    changeset = Context.changeset(context, params)
    # assert 12 == changeset.params.external_id |> String.length()
    assert changeset.valid?
  end

  test "changeset with invalid attributes" do
    changeset = Context.changeset(%Context{}, @invalid_attrs)
    refute changeset.valid?
  end
end
