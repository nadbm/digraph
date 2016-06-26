defmodule Digraffe.LinkControllerTest do
  use Digraffe.ConnCase
  import Digraffe.Factory
  alias Digraffe.Link

  @valid_attrs %{
    title: "some content",
    url: "some content",
  }
  @invalid_attrs %{
    url: -1
  }

  test "lists all entries on index", %{conn: conn} do
    conn = get conn, link_path(conn, :index)
    assert html_response(conn, 200) =~ "Links"
  end

  test "renders form for new resources", %{conn: conn} do
    conn = get conn, link_path(conn, :new)
    assert html_response(conn, 200) =~ "Submit link"
  end

  test "creates resource and redirects when data is valid", %{conn: conn} do
    conn = post conn, link_path(conn, :create), link: @valid_attrs
    assert redirected_to(conn) == link_path(conn, :index)
    assert Repo.get_by(Link, @valid_attrs)
  end

  test "does not create resource and renders errors when data is invalid", %{conn: conn} do
    conn = post conn, link_path(conn, :create), link: @invalid_attrs
    assert html_response(conn, 200) =~ "Add"
  end

  test "shows chosen resource", %{conn: conn} do
    link = create(:link)
    conn = get conn, link_path(conn, :show, link)
    assert html_response(conn, 200) =~ "Back"
  end

  test "renders page not found when id is nonexistent", %{conn: conn} do
    assert_error_sent 404, fn ->
      get conn, link_path(conn, :show, "960c8608-0c58-4dcd-a5f4-13eeb4900292")
    end
  end

  test "renders form for editing chosen resource", %{conn: conn} do
    link = create(:link)
    conn = get conn, link_path(conn, :edit, link)
    assert html_response(conn, 200) =~ "Edit link"
  end

  test "updates chosen resource and redirects when data is valid", %{conn: conn} do
    link = create(:link)
    conn = put conn, link_path(conn, :update, link), link: @valid_attrs
    assert redirected_to(conn) == link_path(conn, :show, link)
    assert Repo.get_by(Link, @valid_attrs)
  end

  test "does not update chosen resource and renders errors when data is invalid", %{conn: conn} do
    link = create(:link)
    conn = put conn, link_path(conn, :update, link), link: @invalid_attrs
    assert html_response(conn, 200) =~ "Edit link"
  end

  test "deletes chosen resource", %{conn: conn} do
    link = create(:link)
    conn = delete conn, link_path(conn, :delete, link)
    assert redirected_to(conn) == link_path(conn, :index)
    refute Repo.get_by(Link, @valid_attrs)
  end
end
