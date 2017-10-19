import EventEmitter from 'events'
import emoji from 'node-emoji'
import mongoose from 'mongoose'

import Love from '../plugins/love'
import Rice from '../plugins/rice'

export default class Core extends EventEmitter {
  
  static defaultOptions = {
    mongo: 'mongodb://localhost:27017/ene'
  }

  constructor (options = {}) {
    super();

    this.options = {
      ...this.constructor.defaultOptions,
      ...options
    }
    this.db = mongoose.connect(this.options.mongo, {
      useMongoClient: true
    })
  }

  reply(str) {
    let toReturn = 'this is a test'

    if(Love.test(str.text)) toReturn = Love.respond(str.text)

    if(Rice.test(str.text)) toReturn = Rice.respond(str.text)

    return emoji.emojify(toReturn)
  }
}
