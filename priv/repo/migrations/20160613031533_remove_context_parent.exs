defmodule Digraffe.Repo.Migrations.RemoveContextParent do
  use Ecto.Migration

  def change do
    alter table(:contexts) do
      remove(:parent_id)
    end
  end
end
