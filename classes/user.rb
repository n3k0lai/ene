
class User

  def constructor(@name)
    @created = Time.now.utc
    @lastspoke = Time.now.utc
  end
end
