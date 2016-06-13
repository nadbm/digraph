defmodule Digraffe.Repo.Migrations.CreateContext do
  use Ecto.Migration

  def change do
    create table(:contexts) do
      add :name, :string
      add :parent_id, references(:contexts, on_delete: :nothing)

      timestamps
    end
    create index(:contexts, [:parent_id])

  end
end
