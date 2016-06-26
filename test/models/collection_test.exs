defmodule Digraffe.CollectionTest do
  use Digraffe.ModelCase

  alias Digraffe.Collection

  @valid_attrs %{title: "some content"}
  @invalid_attrs %{}

  test "changeset with valid attributes" do
    changeset = Collection.changeset(%Collection{}, @valid_attrs)
    assert changeset.valid?
  end

  test "changeset with invalid attributes" do
    changeset = Collection.changeset(%Collection{}, @invalid_attrs)
    refute changeset.valid?
  end
end
