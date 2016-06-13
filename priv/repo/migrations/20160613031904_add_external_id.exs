defmodule Digraffe.Repo.Migrations.AddExternalId do
  use Ecto.Migration

  def change do
    alter table(:contexts) do
      add :external_id, :string
    end
    create unique_index(:contexts, [:external_id])
  end
end
