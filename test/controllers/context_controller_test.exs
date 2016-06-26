defmodule Digraffe.ContextControllerTest do
  use Digraffe.ConnCase
  import Digraffe.Factory
  alias Digraffe.Context

  @valid_attrs %{name: "Home"}
  @invalid_attrs %{name: -1}

  test "lists all entries on index", %{conn: conn} do
    conn = get conn, context_path(conn, :index)
    assert html_response(conn, 200) =~ "Contexts"
  end

  test "renders form for new resources", %{conn: conn} do
    conn = get conn, context_path(conn, :new)
    assert html_response(conn, 200) =~ "New context"
  end

  test "creates resource and redirects when data is valid", %{conn: conn} do
    conn = post conn, context_path(conn, :create), context: @valid_attrs
    assert redirected_to(conn) == context_path(conn, :index)
    assert Repo.get_by(Context, @valid_attrs)
  end

  test "does not create resource and renders errors when data is invalid", %{conn: conn} do
    conn = post conn, context_path(conn, :create), context: @invalid_attrs
    assert html_response(conn, 200) =~ "New context"
  end

  test "shows chosen resource", %{conn: conn} do
    context = create(:context)
    conn = get conn, context_path(conn, :show, context)
    assert html_response(conn, 200) =~ "Show context"
  end

  test "renders page not found when id is nonexistent", %{conn: conn} do
    assert_error_sent 404, fn ->
      get conn, context_path(conn, :show, "960c8608-0c58-4dcd-a5f4-13eeb4900292")
    end
  end

  test "renders form for editing chosen resource", %{conn: conn} do
    context = create(:context)
    conn = get conn, context_path(conn, :edit, context)
    assert html_response(conn, 200) =~ "Edit context"
  end

  test "updates chosen resource and redirects when data is valid", %{conn: conn} do
    context = create(:context)
    conn = put conn, context_path(conn, :update, context), context: @valid_attrs
    assert redirected_to(conn) == context_path(conn, :show, context)
    assert Repo.get_by(Context, @valid_attrs)
  end

  test "does not update chosen resource and renders errors when data is invalid", %{conn: conn} do
    context = create(:context)
    conn = put conn, context_path(conn, :update, context), context: @invalid_attrs
    assert html_response(conn, 200) =~ "Edit context"
  end

  test "deletes chosen resource", %{conn: conn} do
    context = create(:context)
    conn = delete conn, context_path(conn, :delete, context)
    assert redirected_to(conn) == context_path(conn, :index)
    refute Repo.get(Context, context.id)
  end
end
