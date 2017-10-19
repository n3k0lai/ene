<p align="center">
  <img src="./static/images/ene.png" width="350"/>
</p>
<h1 align="center">Ene</h1>
<p align="center">A javascript bot framework</p>

### Summary
Ene is a general bot framework with heavy architectural influence from [munar](https://github.com/welovekpop/munar). Ene looks to provide a robust framework for plugins to interope connections with a variety chat services and social networks.

### Features
Ene is divided into three parts:
* `/src/core/`: The core bot engine
* `/src/adapters/`: Service Input and Output
* `/src/plugins/`: dynamic logic for handling input

### Goals
* Add Rice support
    * Add labelled images
    * Remove labelled images
* Add Games
    * Riichi
        * make rules and validators
        * make simple input and output
        * make simple gameloop
        * add ai
* Add adapters
    * general adapter
    * twitter adapter
    * irc adapter
    * twitch adapter
    * ncurses adapter
    * react adapter
    * discord adapter
