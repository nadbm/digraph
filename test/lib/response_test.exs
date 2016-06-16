defmodule Digraffe.ResponseTest do
  use ExUnit.Case, async: true
  alias Digraffe.Http.Response

  test "#content_type returns the content type from a list of headers" do
    assert {"Content-Type", "text/html"} == Response.find([
      {"Content-Type", "text/html"},
      {"Accept-Language", "en"},
    ], "Content-Type")
  end

  test "#title returns the title of the page from the response" do
    response = %HTTPoison.Response{
      headers: [{"Content-Type", "text/html"}],
      body: "<html><title>Rezrov</title></html>",
    }
    assert "Rezrov" == Response.title(response)
  end

  test "#title returns nil when there is no title" do
    response = %HTTPoison.Response{
      headers: [{"Content-Type", "text/html"}],
      body: "<body><p>hello</p></body>",
    }
    assert nil == Response.title(response)
  end

  test "#title returns nil when the content type is not text/html" do
    response = %HTTPoison.Response{
      headers: [{"Content-Type", "application/json"}],
      body: ~s({"title": "Rezrov"}),
    }
    assert nil == Response.title(response)
  end
end
