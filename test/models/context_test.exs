defmodule Digraffe.ContextTest do
  use Digraffe.ModelCase

  alias Digraffe.Context

  @valid_attrs %{name: "some content"}
  @invalid_attrs %{}

  test "changeset with valid attributes" do
    changeset = Context.changeset(%Context{}, @valid_attrs)
    assert changeset.valid?
  end

  test "changeset with invalid attributes" do
    changeset = Context.changeset(%Context{}, @invalid_attrs)
    refute changeset.valid?
  end
end
