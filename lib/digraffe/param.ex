defimpl Phoenix.Param, for: Digraffe.Link do
  def to_param(%{external_id: id}) do
    "#{id}"
  end
end

defimpl Phoenix.Param, for: Digraffe.Context do
  def to_param(%{external_id: id}) do
    "#{id}"
  end
end

defimpl Phoenix.Param, for: Digraffe.Topic do
  def to_param(%{external_id: id}) do
    "#{id}"
  end
end
