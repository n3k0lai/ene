import config from '../config'
import core from './core'
import T from 'twit'
let tweetQueue = []
const loggingEnabled = true

function getRandomInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

function checkTweetQueue(){
  if (tweetQueue.length > 0){
    var newTweet = tweetQueue.shift()
    if (loggingEnabled === true){
      console.log('Posting new tweet:')
      console.log(newTweet)    
    }

    twitter.post('statuses/update', {
      status: newTweet.text,
      in_reply_to_status_id: newTweet.id
    }, (err, data, response) => {
      if (loggingEnabled === true) {
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
    checkTweetQueue()
  }, 300) //getRandomInt(3000, 60000));
}
const twitter = new T(config.twitter)
console.log('twatter', twitter)
const stream = twitter.stream('statuses/filter', { 
  track: '@' + config.username 
})

stream.on('tweet', (tweet) => {
  const reply = core.reply("local-user", tweet.text)
  if (loggingEnabled) {
    console.log('New Tweet!')
    console.log(tweet.id + ': ' + tweet.text)
    console.log('Adding response to queue:')
    console.log('@' + tweet.user.screen_name + ' ' + reply)
  }

  tweetQueue.push({
    id: tweet.id_str,
    text: '@' + tweet.user.screen_name + ' ' + reply
  })
})

checkTweetQueue()
