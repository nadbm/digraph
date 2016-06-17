defmodule Digraffe.Factory do
  use ExMachina.Ecto, repo: Digraffe.Repo

  alias Digraffe.Util

  def factory(:context) do
    %Digraffe.Context{
      name: "Home",
      external_id: Util.random_id()
    }
  end

  def factory(:topic) do
    %Digraffe.Topic{
      name: "Physics",
      external_id: Util.random_id()
    }
  end

  @url "https://stackoverflow.com/questions/34570612/change-url-to-accept-a-string-instead-id-in-phoenix-framework-elixir"

  def factory(:link) do
    %Digraffe.Link{
      title: "Physics",
      url: @url,
      external_id: Util.id_string(@url)
    }
  end
end
