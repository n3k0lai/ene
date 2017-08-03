
log = (...) -> print "+++", ...

class Twitter extends require "faye.adapter"
  extension_prefix: "faye.plugins."
  new: (@event_loop, @config) =>
  connect: =>
  reconnect: =>
{ :Twitter }
