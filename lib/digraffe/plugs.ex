defmodule Digraffe.Plugs.Authenticate do
  import Plug.Conn
  alias Digraffe.Settings

  def init(options) do
    options
  end

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
