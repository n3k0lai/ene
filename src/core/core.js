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
    super()

    this.options = {
      ...this.constructor.defaultOptions,
      ...options
    }
    this.db = mongoose.connect(this.options.mongo, {
      useMongoClient: true
    })
  }

  receive (adapter, event, args) {
    let toReturn = 'this is a test'

    if (Love.test(event.text)) toReturn = Love.respond(event.text)

    if (Rice.test(event.text)) toReturn = Rice.respond(event.text)

    return emoji.emojify(toReturn)
  }
}
