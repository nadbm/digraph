defmodule Digraffe.ConnCase do
  @moduledoc """
  This module defines the test case to be used by
  tests that require setting up a connection.

  Such tests rely on `Phoenix.ConnTest` and also
  imports other functionality to make it easier
  to build and query models.

  Finally, if the test case interacts with the database,
  it cannot be async. For this reason, every test runs
  inside a transaction which is reset at the beginning
  of the test unless the test case is marked as async.
  """

  use ExUnit.CaseTemplate

  using do
    quote do
      # Import conveniences for testing with connections
      use Phoenix.ConnTest

      alias Digraffe.Repo
      import Ecto
      import Ecto.Changeset
      import Ecto.Query, only: [from: 1, from: 2]

      import Digraffe.Router.Helpers

      # The default endpoint for testing
      @endpoint Digraffe.Endpoint
    end
  end

  setup tags do
    unless tags[:async] do
      Ecto.Adapters.SQL.restart_test_transaction(Digraffe.Repo, [])
    end
    user = %{
      name: "Gnusto",
      provider: "github",
      provider_id: "https://api.github.com/users/gnusto",
      avatar_url: "https://avatars.githubusercontent.com/u/760949?v=3"
    }
    conn = Phoenix.ConnTest.conn()
    |> Plug.Conn.put_private(:current_user, user)
    {:ok, conn: conn}
  end
end
