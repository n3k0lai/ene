import Core from '.core/core'
import config from '../config'
import TwitterAdapter from './adapters/twitter'

let Ene = new Core(config);
let Twitter = new TwitterAdapter(Ene, config);

Twitter.connect();
