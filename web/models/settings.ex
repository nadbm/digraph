defmodule Digraffe.Settings do
  use Digraffe.Web, :model

  alias Digraffe.{Repo, Settings}

  schema "settings" do
    field :name,        :string
    field :provider,    :string
    field :provider_id, :string
    field :avatar_url,  :string
    has_many :collections,         Digraffe.Collection, foreign_key: :owner_id
    has_one  :selected_collection, Digraffe.Collection, foreign_key: :id
    timestamps
  end

  @required_fields ~w(name provider provider_id avatar_url)
  @optional_fields ~w()

  def changeset(model, params \\ :empty) do
    model
    |> cast(params, @required_fields, @optional_fields)
    |> unique_constraint(:provider_id)
  end

  def get_or_create(user) when is_map(user) do
    case get_settings(user) do
      nil   -> changeset(%Settings{}, user) |> Repo.insert!
      model -> model
    end
  end

  def get_or_create(_) do
    nil
  end

  defp get_settings(%{} = user) do
    case Repo.get_by(Settings, user) do
      nil      -> nil
      settings -> Repo.preload(settings, :selected_collection)
    end
  end

  def small_avatar_url(settings) do
    ~s(#{settings.avatar_url}&s=40)
  end
end
