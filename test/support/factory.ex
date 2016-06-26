defmodule Digraffe.Factory do
  use ExMachina.Ecto, repo: Digraffe.Repo

  def factory(:context) do
    %Digraffe.Context{
      name: "Home",
    }
  end

  def factory(:topic) do
    %Digraffe.Topic{
      name: "Physics",
    }
  end

  @url "https://stackoverflow.com/questions/34570612/change-url-to-accept-a-string-instead-id-in-phoenix-framework-elixir"

  def factory(:link) do
    %Digraffe.Link{
      title: "Physics",
      url: @url,
    }
  end
end
