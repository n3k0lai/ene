import EventEmitter from 'events'
import T from 'twit'

export default class Twitter {
  static loggingEnabled = true
  
  static getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }

  constructor (bot, options) {
    this.tweetQueue = []
    this.bot = bot
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
        const reply = this.bot.reply(tweet)
        if (this.loggingEnabled) {
          console.log('New Tweet!')
          console.log(tweet.id + ': ' + tweet.text)
          console.log('Adding response to queue:')
          console.log('@' + tweet.user.screen_name + ' ' + reply)
        }

        this.tweetQueue.push({
          id: tweet.id_str,
          text: '@' + tweet.user.screen_name + ' ' + reply
        })
      }) 
      this.checkTweetQueue()
    })
  }

  checkTweetQueue(){
    if (this.tweetQueue.length > 0){
      let newTweet = this.tweetQueue.shift()
      if (this.loggingEnabled === true){
        console.log('Posting new tweet:')
        console.log(newTweet)    
      }

      this.client.post('statuses/update', {
        status: newTweet.text,
        in_reply_to_status_id: newTweet.id
      }, (err, data, response) => {
        if (this.loggingEnabled === true) {
          if (err) {
            console.log('ERROR')
            console.log(err)       
          } else {
            console.log('NO ERROR')
          }
        }
      })
    }

    setTimeout(() => {
      this.checkTweetQueue()
    }, 300) //getRandomInt(3000, 60000));
  }
}
