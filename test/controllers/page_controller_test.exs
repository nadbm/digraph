defmodule Digraffe.PageControllerTest do
  use Digraffe.ConnCase

  test "GET /", %{conn: conn} do
    conn = get conn, "/"
    assert redirected_to(conn) == link_path(conn, :index)
  end
end
