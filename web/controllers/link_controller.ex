defmodule Digraffe.LinkController do
  use Digraffe.Web, :controller

  alias Digraffe.Link

  plug :scrub_params, "link" when action in [:create, :update]

  def index(conn, _params) do
    render(conn, "index.html",
      links:     Repo.all(Link),
      changeset: Link.changeset(%Link{}))
  end

  def new(conn, _params) do
    changeset = Link.changeset(%Link{})
    render(conn, "new.html", changeset: changeset)
  end

  def create(conn, %{"link" => params}) do
    params = Link.params_for_create(params)
    case params do

      nil ->
        conn
        |> put_flash(:info, "There was a problem.")
        |> redirect(to: link_path(conn, :index))

      _ ->
        changeset = Link.changeset(%Link{}, params)
        case Repo.insert(changeset) do

          {:ok, _link} ->
            conn
            |> put_flash(:info, "Link added.")
            |> redirect(to: link_path(conn, :index))

          {:error, changeset} ->
            case changeset do
              %{errors: [external_id: "has already been taken"]} ->
                conn
                |> put_flash(:info, "Link found.")
                |> redirect(to: link_path(conn, :index))
              _ ->
                render(conn, "index.html",
                  changeset: changeset,
                  links: Repo.all(Link))
            end
        end
    end
  end

  def show(conn, %{"external_id" => id}) do
    link = Repo.get_by!(Link, external_id: id)
    render(conn, "show.html", link: link)
  end

  def edit(conn, %{"external_id" => id}) do
    link = Repo.get_by!(Link, external_id: id)
    changeset = Link.changeset(link)
    render(conn, "edit.html", link: link, changeset: changeset)
  end

  def update(conn, %{"external_id" => id, "link" => params}) do
    link = Repo.get_by!(Link, external_id: id)
    changeset = Link.changeset(link, params)
    case Repo.update(changeset) do
      {:ok, link} ->
        conn
        |> put_flash(:info, "Link updated.")
        |> redirect(to: link_path(conn, :show, link))
      {:error, changeset} ->
        render(conn, "edit.html", link: link, changeset: changeset)
    end
  end

  def delete(conn, %{"external_id" => id}) do
    link = Repo.get_by!(Link, external_id: id)
    # Here we use delete! (with a bang) because we expect
    # it to always work (and if it does not, it will raise).
    Repo.delete!(link)
    conn
    |> put_flash(:info, "Link deleted successfully.")
    |> redirect(to: link_path(conn, :index))
  end
end
