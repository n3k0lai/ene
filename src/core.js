import emoji from 'node-emoji'

import Love from './love'
import Rice from './rice'

module.exports = {
  reply(str) {
    let toReturn = 'this is a test'

    if(Love.test(str.text)) toReturn = Love.respond(str.text)

    if(Rice.test(str.text)) toReturn = Rice.respond(str.text)

    return emoji.emojify(toReturn)
  }
}
