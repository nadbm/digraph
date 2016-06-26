defmodule Digraffe.Repo.Migrations.AddCollectionOwner do
  use Ecto.Migration

  def change do
    alter table(:collections) do
      add :owner_id, references(:settings, type: :uuid)
    end
  end
end
