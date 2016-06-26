defmodule Digraffe.Settings do
  defstruct [:name, :provider, :avatar_url, :small_avatar_url]

  def new(settings) when is_map(settings) do
    settings
    |> Map.merge(%{small_avatar_url: ~s(#{settings.avatar_url}&s=40)})
  end

  def new(_) do
    nil
  end
end
