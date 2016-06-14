defmodule Digraffe.ContextTest do
  use Digraffe.ModelCase
  import Digraffe.Factory
  alias Digraffe.{Context, Util}

  @valid_attrs %{name: "Home"}
  @invalid_attrs %{}

  test "changeset with valid attributes" do
    context = create(:context)
    params = Util.params_for_create(@valid_attrs)
    changeset = Context.changeset(context, params)
    assert changeset.valid?
  end

  test "changeset with invalid attributes" do
    changeset = Context.changeset(%Context{}, @invalid_attrs)
    refute changeset.valid?
  end
end
