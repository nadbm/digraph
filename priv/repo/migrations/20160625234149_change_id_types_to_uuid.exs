defmodule Digraffe.Repo.Migrations.ChangeIdTypesToUuid do
  use Ecto.Migration

  def change do
    alter table(:links, primary_key: false) do
      remove :id
      add :id, :uuid, primary_key: true, autogenerate: true
      remove :external_id
    end

    alter table(:contexts, primary_key: false) do
      remove :id
      add :id, :uuid, primary_key: true, autogenerate: true
      remove :external_id
    end

    alter table(:topics, primary_key: false) do
      remove :id
      add :id, :uuid, primary_key: true, autogenerate: true
      remove :external_id
    end
  end

  def down do
    alter table(:links, primary_key: true) do
      remove :id
      add :id, :integer, primary_key: true, autogenerate: true
      add :external_id, :string
    end

    alter table(:contexts, primary_key: true) do
      remove :id
      add :id, :integer, primary_key: true, autogenerate: true
      add :external_id, :string
    end

    alter table(:topics, primary_key: true) do
      remove :id
      add :id, :integer, primary_key: true, autogenerate: true
      add :external_id, :string
    end
  end
end
