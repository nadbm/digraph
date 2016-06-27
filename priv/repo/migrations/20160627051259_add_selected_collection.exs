defmodule Digraffe.Repo.Migrations.AddSelectedCollection do
  use Ecto.Migration

  def change do
    alter table(:settings) do
      add :selected_collection_id, references(:collections, type: :uuid)
    end
  end
end
