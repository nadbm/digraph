defmodule Digraffe.ContextTest do
  use Digraffe.ModelCase
  import Digraffe.Factory
  alias Digraffe.Context

  @valid_attrs %{name: "Home"}
  @invalid_attrs %{}

  test "changeset with valid attributes" do
    context = create(:context)
    changeset = Context.changeset(context, @valid_attrs)
    assert changeset.valid?
  end

  test "changeset with invalid attributes" do
    changeset = Context.changeset(%Context{}, @invalid_attrs)
    refute changeset.valid?
  end
end
