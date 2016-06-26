defmodule Digraffe.Router do
  use Digraffe.Web, :router
  alias Digraffe.Settings

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_flash
    plug :protect_from_forgery
    plug :put_secure_browser_headers
    plug :assign_settings
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/", Digraffe do
    pipe_through :browser # Use the default browser stack

    get "/", PageController, :index

    resources "/contexts",    ContextController
    resources "/topics",      TopicController
    resources "/links",       LinkController
    resources "/collections", CollectionController
  end

  scope "/auth", Digraffe do
    pipe_through :browser

    get "/:provider",          AuthController, :index
    get "/:provider/callback", AuthController, :callback
    delete "/logout",          AuthController, :delete
  end

  # Fetch the current user from the session and add it to `conn.assigns`. This
  # will allow you to have access to the current user in your views with
  # `@current_account`.
  defp assign_settings(conn, _) do
    user = get_session(conn, :current_user)
    assign conn, :settings, Settings.get_or_create(user)
  end
end
