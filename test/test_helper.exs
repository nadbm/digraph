ExUnit.start

Mix.Task.run "ecto.create", ~w(-r Digraffe.Repo --quiet)
Mix.Task.run "ecto.migrate", ~w(-r Digraffe.Repo --quiet)
Ecto.Adapters.SQL.begin_test_transaction(Digraffe.Repo)

