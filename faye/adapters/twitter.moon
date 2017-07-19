
socket = require "socket"
{parse: parse_url} = require "socket.url"

import decode_html_entities from require "faye.util"
import insert from table

import Reader from require "faye.socket"
import Dispatch from require "faye.dispatch"

log = (...) -> print "+++", ...

class Twitter extends require "faye.adapter"
  extension_prefix: "faye.plugins."

  new: (@event_loop, @config) =>
    host = assert @config.host, "config missing host"
    @host, @port = host\match "^(.-):(%d*)$"

    @host or= @config.host
    @port or= 6667

    @dispatch = Dispatch!

    @extensions = for e in *@config.extensions or {}
      require("#{@extension_prefix}#{e}") @

    @dispatch\trigger "irc.before_connect", @

    @connect!

    irc = @
    config = @config
    @reader = Reader @socket, {
      loop: =>
        while true
          irc\handle_message @get_line!

      handle_error: (msg) =>
        irc.socket = nil
        if msg == "closed"
          log "Disconnected. Reconnecting in #{config.reconnect_time} seconds"
          irc\reconnect!
        else
          error msg
    }

    @event_loop\add_listener @reader

  connect: =>
    @channels = {}
    socket = socket.connect @host, @port
    @socket = socket

    unless @socket
      error "could not connect to server #{@host}:#{@port}"

    if @config.oauth_token
      @socket\send "PASS #{@config.oauth_token}\n\n"

    @socket\send "NICK #{@config.name}\r\n"
    @socket\send "USER ".."moon "\rep(3)..":Bildo Bagins\r\n"

    @dispatch\trigger "irc.connect", @

    @event_loop\add_task {
      time: @config.join_delay or 1
      action: ->
        return unless @socket

        if @config.password
          @message_to "NickServ", "IDENTIFY #{@config.password}"

        for channel in *@config.channels
          @join channel
    }

  reconnect: =>
    @event_loop\add_task {
      interval: @config.reconnect_time
      action: (task) ->
        log "Reconnected:", pcall ->
          log "Trying to reconnect"
          @connect!
          @reader\set_socket @socket
          @event_loop\add_listener @reader
          task.interval = nil -- stop trying to reconnect
    }

{ :Twitter }
