defmodule Digraffe.CollectionView do
  use Digraffe.Web, :view

  def selected_collection(settings, collection) do
    if collection_is_selected(settings, collection), do: "active", else: ""
  end

  defp collection_is_selected(settings, collection) do
    case settings.selected_collection do
      nil -> false
      id  -> id == collection.id
    end
  end
end
