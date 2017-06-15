config = require "lapis.config"

config "test", ->
  postgres {
    database: "fayebot_test"
  }

config "development", ->
  join_delay 2

  host "localhost"
  name "faye"

  channels { "#fayebot" }

  admin_password "admin"

  extensions {
    "url_titles"
    "admin"
    "stats"
  }

  postgres {
    database: "faye"
  }

  systemd {
    user: true
  }


