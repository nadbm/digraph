defmodule Digraffe.Repo.Migrations.CreateTopic do
  use Ecto.Migration

  def change do
    create table(:topics) do
      add :name, :string
      add :external_id, :string
      timestamps
    end
  end
end
