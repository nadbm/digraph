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

  test "#normalize_url provides a canonical version of a url" do
    assert {:ok, "https://rezrov.org/"} == Util.normalize_url("https://rezrov.org")
  end

  test "#normalize_url fills in the schema" do
    assert {:ok, "http://rezrov.org/"} == Util.normalize_url("//rezrov.org")
  end

  test "#normalize_url lowercases the string" do
    assert {:ok, "http://rezrov.org/"} == Util.normalize_url("Http://Rezrov.Org")
  end

  test "#normalize_url deals with bad input" do
    assert {:error, "http//rezrov.org"} == Util.normalize_url("http//rezrov.org")
    assert {:error, "http://rezrovorg"} == Util.normalize_url("http://rezrovorg")
  end
end
