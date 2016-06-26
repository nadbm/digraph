defmodule Digraffe.SettingsTest do
  use ExUnit.Case, async: true

  alias Digraffe.Settings

  test "#small_avatar_url is added" do
    settings = Settings.new(%{avatar_url: "http://avatar/url?v=1"})
    assert "http://avatar/url?v=1&s=40" == settings.small_avatar_url
  end
end
