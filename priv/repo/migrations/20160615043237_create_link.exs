defmodule Digraffe.Repo.Migrations.CreateLink do
  use Ecto.Migration

  def change do
    create table(:links) do
      add :title, :string
      add :url, :string
      add :external_id, :string
      timestamps
    end
  end
end
