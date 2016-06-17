defmodule Digraffe.TopicController do
  use Digraffe.Web, :controller

  alias Digraffe.{Topic, Util}

  plug :scrub_params, "topic" when action in [:create, :update]

  def index(conn, _params) do
    topics = Repo.all(Topic)
    render(conn, "index.html", topics: topics)
  end

  def new(conn, _params) do
    changeset = Topic.changeset(%Topic{})
    render(conn, "new.html", changeset: changeset)
  end

  def create(conn, %{"topic" => params}) do
    params = Util.params_for_create(params)
    changeset = Topic.changeset(%Topic{}, params)
    case Repo.insert(changeset) do
      {:ok, _topic} ->
        conn
        |> put_flash(:info, "Topic created successfully.")
        |> redirect(to: topic_path(conn, :index))
      {:error, changeset} ->
        render(conn, "new.html", changeset: changeset)
    end
  end

  def show(conn, %{"external_id" => id}) do
    topic = Repo.get_by!(Topic, external_id: id)
    render(conn, "show.html", topic: topic)
  end

  def edit(conn, %{"external_id" => id}) do
    topic = Repo.get_by!(Topic, external_id: id)
    changeset = Topic.changeset(topic)
    render(conn, "edit.html", topic: topic, changeset: changeset)
  end

  def update(conn, %{"external_id" => id, "topic" => params}) do
    topic = Repo.get_by!(Topic, external_id: id)
    changeset = Topic.changeset(topic, params)
    case Repo.update(changeset) do
      {:ok, topic} ->
        conn
        |> put_flash(:info, "Topic updated successfully.")
        |> redirect(to: topic_path(conn, :show, topic))
      {:error, changeset} ->
        render(conn, "edit.html", topic: topic, changeset: changeset)
    end
  end

  def delete(conn, %{"external_id" => id}) do
    topic = Repo.get_by!(Topic, external_id: id)
    # Here we use delete! (with a bang) because we expect
    # it to always work (and if it does not, it will raise).
    Repo.delete!(topic)
    conn
    |> put_flash(:info, "Topic deleted successfully.")
    |> redirect(to: topic_path(conn, :index))
  end
end
