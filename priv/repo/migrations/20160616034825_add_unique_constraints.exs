defmodule Digraffe.Repo.Migrations.AddUniqueConstraints do
  use Ecto.Migration

  def change do
    create unique_index(:topics, [:external_id])
    create unique_index(:links,  [:external_id])
  end
end
