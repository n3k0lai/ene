twitter = require "luatwit"

log = (...) -> print "+++", ...

class Twitter extends require "faye.adapter"
  extension_prefix: "faye.plugins."
  new: (@event_loop, @config) =>

  connect: =>
    @client = twitter.api.new(twitter.util.loadkeys(@config.keys.app_keys, @config.keys.user_keys))

  reconnect: =>

{ :Twitter }
