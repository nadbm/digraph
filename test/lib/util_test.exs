defmodule Digraffe.UtilTest do
  use ExUnit.Case, async: true
  alias Digraffe.Util

  test "#simple_chars removes hyphens from a string" do
    assert "123" == Util.simple_chars("-_=123-_=")
  end

  test "#string_keys converts symbol keys to strings" do
    assert %{"k1" => 1} == Util.string_keys(%{k1: 1})
    assert %{"k2" => 2} == Util.string_keys(%{"k2" => 2})
    assert %{"k1" => 1, "k2" => 2} == Util.string_keys(%{:k1 => 1, "k2" => 2})
  end

  test "#params_for_create adds an external_id" do
    params = Util.params_for_create(%{name: "Home"})
    assert Map.has_key?(params, "external_id")
    assert 12 == String.length(params["external_id"])
  end
end
