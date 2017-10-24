import Core from './core/core'
import config from '../config'
import TwitterAdapter from './adapters/twitter'
//import IrcAdapter from './adapters/irc'
import DiscordAdapter from './adapters/discord'

let Ene = new Core(config);
let Twitter = new TwitterAdapter(Ene, config.twitter);
//let Irc = new IrcAdapter(Ene, config.irc)
let Discord = new DiscordAdapter(Ene, config.discord);

Twitter.connect();
