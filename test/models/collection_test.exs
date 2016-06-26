defmodule Digraffe.CollectionTest do
  use Digraffe.ModelCase
  alias Digraffe.Collection

  @valid_attrs %{
    title: "Some collection",
    owner_id: "d11d32c6-f787-4178-aa6a-bacd9002acd1"
  }
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
