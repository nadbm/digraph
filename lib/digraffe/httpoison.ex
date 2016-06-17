defmodule Digraffe.Httpoison.Fake do

  def get(_string) do
    { :ok,
      %HTTPoison.Response{
        headers: [{"Content-Type", "text/http"}],
        body: "<html><title>Gnusto</title></html>"
      }
    }
  end
end
