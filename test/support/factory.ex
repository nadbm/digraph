defmodule Digraffe.Factory do
  use ExMachina.Ecto, repo: Digraffe.Repo

  def factory(:context) do
    %Digraffe.Context{
      name: "Home"
    }
  end
end
