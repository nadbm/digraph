defmodule Digraffe.TopicTest do
  use Digraffe.ModelCase
  import Digraffe.Factory
  alias Digraffe.{Topic, Util}

  @valid_attrs %{external_id: "some content", name: "some content"}
  @invalid_attrs %{}

  test "changeset with valid attributes" do
    topic = create(:topic)
    params = Util.params_for_create(@valid_attrs)
    changeset = Topic.changeset(topic, params)
    assert changeset.valid?
  end

  test "changeset with invalid attributes" do
    changeset = Topic.changeset(%Topic{}, @invalid_attrs)
    refute changeset.valid?
  end
end
