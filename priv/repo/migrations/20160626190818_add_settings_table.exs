defmodule Digraffe.Repo.Migrations.AddSettingsTable do
  use Ecto.Migration

  def change do
    create table(:settings, primary_key: false) do
      add :id, :uuid, primary_key: true
      add :name, :string
      add :provider, :string
      add :provider_id, :string
      add :avatar_url, :string
      timestamps
    end
    create unique_index(:settings, [:provider_id])
  end
end
