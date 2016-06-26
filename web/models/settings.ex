defmodule Digraffe.Settings do
  use Digraffe.Web, :model

  alias Digraffe.{Repo, Settings}

  schema "settings" do
    field :name, :string
    field :provider, :string
    field :provider_id, :string
    field :avatar_url, :string
    has_many :collections, Digraffe.Collection, foreign_key: :owner_id
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
    case model = Repo.get_by(Settings, user) do
      nil -> changeset(%Settings{}, user) |> Repo.insert!
      _   -> model
    end
  end

  def get_or_create(_) do
    nil
  end

  def small_avatar_url(settings) do
    ~s(#{settings.avatar_url}&s=40)
  end
end
