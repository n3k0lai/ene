import EventLoop from require "faye.event_loop"
import Irc from require "faye.adapter.irc"

loop = EventLoop!
irc = Irc loop, require("faye.config")
loop\run!

