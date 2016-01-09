
class User
  include CouchPotato::Persistence

  poperty :name, :type => String
  validates_presence_of :name
  property :created, :type => Date, :default => Proc.new { Time.now.utc }
  property :lastSpoke, :type => Date, :default => Proc.new { Time.now.utc }
  property :randomPhrase, :type => String, :default => ''
  property :wallet, :type => Fixnum, :default => 0
  property :pseudonym, :type => [String], :default => []
  property :waifu, :type => String, :default => ''
  property :husbando, :type => String, :default => ''
  property :selfie, :type => String, :default => ''

end
