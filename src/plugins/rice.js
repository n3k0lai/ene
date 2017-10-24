import Plugin from '../core/plugin.js'

export default class Rice extends Plugin {
  constructor() {

  }

  static test(str) {
    return str.includes('add')
  }

  static respond(str) {
    return 'added'
  }
}
