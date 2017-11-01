import Eris from 'eris'

import Adapter from '../core/adapter'

export default class Discord extends Adapter {
  static adapterName = 'Discord'

  constructor (bot, options) {
    super(bot)
    this.options = options
  }

  async connect () {
    return new Promise((resolve, reject) => {
      this.client = new Eris(this.options.token)
      console.log('Discord :>', this.client)

      this.client.on('messageCreate', (msg) => {
        let reply = this.receive(msg.content)

        console.log('New discord msg :>')
        console.log('@' + msg.member.username + ': ' + msg.content)
        console.log('Adding response to queue:')
        console.log(reply)

        this.client.createMessage(msg.channel.id, reply)
      })

      this.client.connect()
    })
  }

  disconnect () {
    this.client.disconnect()
    this.client = null
  }

  getSelf () {
    return this.self
  }
}
