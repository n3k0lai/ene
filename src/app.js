import Core from './core'
import config from '../config'
import TwitterAdapter from './adapters/twitter'

let Ene = new Core(config);
let Twitter = new TwitterAdapter(Ene, config);

Twitter.connect();
