defmodule Digraffe.Repo.Migrations.CreateCollection do
  use Ecto.Migration

  def change do
    create table(:collections, primary_key: false) do
      add :id, :uuid, primary_key: true
      add :title, :string
      timestamps
    end
  end
end
