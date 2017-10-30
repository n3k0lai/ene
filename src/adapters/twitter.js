import T from 'twit'

import Adapter from '../core/adapter'

export default class Twitter extends Adapter {
  static adapterName = 'twitter'

  constructor (bot, options) {
    super(bot)
    this.tweetQueue = []
    this.options = options
  }

  async connect () {
    return new Promise((resolve, reject) => {
      this.client = new T(this.options.creds)
      console.log('twatter', this.client)

      this.stream = this.client.stream('statuses/filter', {
        track: '@' + this.options.username
      })

      this.stream.on('tweet', (tweet) => {
        const reply = this.receive(tweet)

        console.log('New Tweet!')
        console.log(tweet.id + ': ' + tweet.text)
        console.log('Adding response to queue:')
        console.log('@' + tweet.user.screen_name + ' ' + reply)

        this.tweetQueue.push({
          id: tweet.id_str,
          text: '@' + tweet.user.screen_name + ' ' + reply
        })
      })
      this.checkTweetQueue()
    })
  }

  disconnect () {
    this.stream.stop()
    this.stream = null
    this.client = null
    this.tweetQueue = []
  }

  getSelf () {
    return this.self
  }

  checkTweetQueue () {
    if (this.tweetQueue.length > 0) {
      let newTweet = this.tweetQueue.shift()

      console.log('Posting new tweet:')
      console.log(newTweet)

      this.client.post('statuses/update', {
        status: newTweet.text,
        in_reply_to_status_id: newTweet.id
      }, (err, data, response) => {
        if (err) {
          console.log('ERROR')
          console.log(err)
        } else {
          console.log('NO ERROR')
        }
      })
    }

    setTimeout(() => {
      this.checkTweetQueue()
    }, 300)
  }
}
