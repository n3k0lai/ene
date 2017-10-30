import Core from './core/core'
import config from '../config'
import TwitterAdapter from './adapters/twitter'
// import IrcAdapter from './adapters/irc'

let Ene = new Core(config)
let Twitter = new TwitterAdapter(Ene, config.twitter)
// let Irc = new IrcAdapter(Ene, config.irc)

Twitter.connect()
