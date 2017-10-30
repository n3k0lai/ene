import Plugin from '../core/plugin.js'

<<<<<<< HEAD
export default class Rice {
  static test (str) {
=======
export default class Rice extends Plugin {
  constructor() {

  }

  static test(str) {
>>>>>>> 7ad87cb26ccf7d4bfb4ea0a78bfa407aa395fd42
    return str.includes('add')
  }

  static respond (str) {
    return 'added'
  }
}
