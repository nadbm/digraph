defmodule Digraffe.Router do
  use Digraffe.Web, :router

  pipeline :browser do
    plug :accepts, ["html"]
    plug :fetch_session
    plug :fetch_flash
    plug :protect_from_forgery
    plug :put_secure_browser_headers
  end

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/", Digraffe do
    pipe_through :browser # Use the default browser stack

    get "/", PageController, :index

    resources "/contexts", ContextController, param: "external_id"
    resources "/topics",   TopicController,   param: "external_id"
    resources "/links",    LinkController,    param: "external_id"
  end
end
