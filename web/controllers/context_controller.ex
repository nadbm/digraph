defmodule Digraffe.ContextController do
  use Digraffe.Web, :controller

  alias Digraffe.{Context, Util}

  plug :scrub_params, "context" when action in [:create, :update]

  def index(conn, _params) do
    contexts = Repo.all(Context)
    render(conn, "index.html", contexts: contexts)
  end

  def new(conn, _params) do
    changeset = Context.changeset(%Context{})
    render(conn, "new.html", changeset: changeset)
  end

  def create(conn, %{"context" => params}) do
    params = Util.params_for_create(params)
    changeset = Context.changeset(%Context{}, params)
    case Repo.insert(changeset) do
      {:ok, _context} ->
        conn
        |> put_flash(:info, "Context created successfully.")
        |> redirect(to: context_path(conn, :index))
      {:error, changeset} ->
        render(conn, "new.html", changeset: changeset)
    end
  end

  def show(conn, %{"external_id" => id}) do
    context = Repo.get_by!(Context, external_id: id)
    render(conn, "show.html", context: context)
  end

  def edit(conn, %{"external_id" => id}) do
    context = Repo.get_by!(Context, external_id: id)
    changeset = Context.changeset(context)
    render(conn, "edit.html", context: context, changeset: changeset)
  end

  def update(conn, %{"external_id" => id, "context" => params}) do
    context = Repo.get_by!(Context, external_id: id)
    changeset = Context.changeset(context, params)

    case Repo.update(changeset) do
      {:ok, context} ->
        conn
        |> put_flash(:info, "Context updated successfully.")
        |> redirect(to: context_path(conn, :show, context))
      {:error, changeset} ->
        render(conn, "edit.html", context: context, changeset: changeset)
    end
  end

  def delete(conn, %{"external_id" => id}) do
    context = Repo.get_by!(Context, external_id: id)

    # Here we use delete! (with a bang) because we expect
    # it to always work (and if it does not, it will raise).
    Repo.delete!(context)

    conn
    |> put_flash(:info, "Context deleted successfully.")
    |> redirect(to: context_path(conn, :index))
  end
end
