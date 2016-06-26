defmodule Digraffe.Link do
  use Digraffe.Web, :model

  alias Digraffe.Http

  @http_client Application.get_env(:digraffe, :http_client)

  schema "links" do
    field :title, :string
    field :url,   :string
    timestamps
  end

  @required_fields ~w(title url)
  @optional_fields ~w()

  def changeset(model, params \\ :empty) do
    params = add_title(params)
    model
    |> cast(params, @required_fields, @optional_fields)
  end

  defp add_title(%{"url" => url} = params) do
    case @http_client.get(url) do
      {:ok, response} ->
        unless title = Http.Response.title(response) do
          title = url
        end
        Map.put(params, "title", title)
      _ ->
        params
    end
  end

  defp add_title(params) do
    params
  end
end
