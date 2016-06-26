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

  def factory(:settings) do
    %Digraffe.Settings{
      name: "Rezrov",
      provider: "github",
      provider_id: "https://api.github.com/users/rezrov",
      avatar_url: "https://avatars.githubusercontent.com/u/760949?v=3"
    }
  end

  def factory(:collection) do
    owner = create(:settings)
    %Digraffe.Collection{
      title: "Links",
      owner_id: owner.id
    }
  end
end
