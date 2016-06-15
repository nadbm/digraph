defmodule Digraffe.LinkTest do
  use Digraffe.ModelCase

  alias Digraffe.Link

  @valid_attrs %{
    title: "Some interesting blog post",
    url: "http://google.com/blog",
    external_id: "123456789ab"
  }
  @invalid_attrs %{}

  test "changeset with valid attributes" do
    params = Link.params_for_create(@valid_attrs)
    changeset = Link.changeset(%Link{}, params)
    assert changeset.valid?
  end

  test "changeset with invalid attributes" do
    changeset = Link.changeset(%Link{}, @invalid_attrs)
    refute changeset.valid?
  end
end
