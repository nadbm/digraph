defmodule Digraffe.Plugs.Authenticate do
  import Plug.Conn
  alias Digraffe.Settings

  def init(options), do: options

  def call(conn, _) do
    user = current_user(conn)
    assign(conn, :settings, Settings.get_or_create(user))
  end

  defp current_user(conn) do
    case Mix.env do
      :test ->
        conn.private[:current_user]
      _ ->
        conn |> fetch_session |> get_session(:current_user)
    end
  end
end

defmodule Digraffe.Plugs.SetCollection do
  import Plug.Conn
  import Digraffe.Router.Helpers
  import Phoenix.Controller

  def init(options), do: options

  def call(conn, _assigns) do
    case conn.assigns.settings.selected_collection do
      nil ->
        conn
        |> put_flash(:info, "A collection must be selected")
        |> redirect(to: collection_path(conn, :index))
        |> halt
      _collection ->
        conn
    end
  end
end
