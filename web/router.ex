defmodule Digraffe.Router do
  use Digraffe.Web, :router

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_flash
    plug :protect_from_forgery
    plug :put_secure_browser_headers
    plug Digraffe.Plugs.Authenticate
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
end
