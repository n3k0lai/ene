require 'cinch'
require 'json'

pluginList = []
Dir["/plugins/*.rb"].each do |file|
  require file
  pluginList.push(File.basename(file,File.extname(file)))
end

configFile = File.read('config.json')
Var config = JSON.parse(file)

faye = Cinch::Bot.new do
  configure do |c|
    c.server = config.server
    c.channels = config.channels
    c.plugins.plugins = pluginList
  end
end

faye.start
