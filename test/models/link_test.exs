defmodule Digraffe.LinkTest do
  use Digraffe.ModelCase

  alias Digraffe.Link

  @valid_attrs %{
    title: "Some interesting blog post",
    url: "http://google.com/blog",
  }
  @invalid_attrs %{}

  test "changeset with valid attributes" do
    changeset = Link.changeset(%Link{}, @valid_attrs)
    assert changeset.valid?
  end

  test "changeset with invalid attributes" do
    changeset = Link.changeset(%Link{}, @invalid_attrs)
    refute changeset.valid?
  end
end
