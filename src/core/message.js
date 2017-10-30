
export default class Message {
  constructor (source, text, sourceMessage) {
    this.source = source
    this.text = text
    this.sourceMessage = sourceMessage
  }

  toString () {
    return `${this.source} | <${this.username || this.user || 'Unknown'}> ${this.text}`
  }
}
