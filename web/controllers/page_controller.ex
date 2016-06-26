defmodule Digraffe.PageController do
  use Digraffe.Web, :controller

  def index(conn, _params) do
    redirect conn, to: link_path(conn, :index)
  end
end
