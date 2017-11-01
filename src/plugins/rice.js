import Plugin from '../core/plugin.js'

export default class Rice extends Plugin {

  static test (str) {
    return str.includes('add')
  }

  static respond (str) {
    return 'added'
  }

  static testAdd (str) {
    return str.includes('add')
  }

  static testRm (str) {
    return str.includes('rm') || str.includes('rem') || str.includes('del')
  }

  static testHome (str) {
    return str.includes('home') || str.includes('homescreen') || str.includes('homescreens')
  }

  static testDtop (str) {
    return str.includes('desk') || str.includes('desktop') || str.includes('dtop')
  }

  static testWaifu (str) {
    return str.includes('waifu')
  }

  static testHusbando (str) {
    return str.includes('husbando')
  }

  static testBattleStation (str) {
    return str.includes('battlestation') || str.includes('bullshit') || str.includes('bs')
  }

  static testMemer (str) {
    return str.includes(':^)')
  }

  static testQ (str) {
    return str.includes('??')
  }

  static testNp (str) {
    return str.includes('np') || str.includes('playing')
  }
}
