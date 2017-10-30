import Core from './core/core'
import config from '../config'
import TwitterAdapter from './adapters/twitter'
<<<<<<< HEAD
// import IrcAdapter from './adapters/irc'

let Ene = new Core(config)
let Twitter = new TwitterAdapter(Ene, config.twitter)
// let Irc = new IrcAdapter(Ene, config.irc)
=======
//import IrcAdapter from './adapters/irc'
import DiscordAdapter from './adapters/discord'

let Ene = new Core(config);
let Twitter = new TwitterAdapter(Ene, config.twitter);
//let Irc = new IrcAdapter(Ene, config.irc)
let Discord = new DiscordAdapter(Ene, config.discord);
>>>>>>> 7ad87cb26ccf7d4bfb4ea0a78bfa407aa395fd42

Twitter.connect()
