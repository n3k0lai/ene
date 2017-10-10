import ChannelUsers from require "faye.models"

class Stats extends require  "faye.plugin"
  new: (@irc) =>
    @irc\on "irc.message", @\message_handler

  message_handler: (e, irc, name, channel, message) =>
    return unless channel\match "^#"
    ChannelUsers\log channel, name, message

