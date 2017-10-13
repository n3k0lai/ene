<p align="center">
  <img src="./static/images/ene.png" width="350"/>
</p>
<h1 align="center">Ene</h1>
<p align="center">A javascript bot framework</p>

### Summary
Ene is a general bot framework originally forked from [leafo's bot](https://github.com/leafo/saltw-bot) with architectural influence from [munar](https://github.com/welovekpop/munar). Ene looks to provide a robust framework for plugins to interope connections with a variety of other chat services. The main intention for Ene is a strong AI for Riichi Mahjong because I want another option from Tenhou.

### Features
Ene is divided into three parts:
* `Ene`: The core game engine
* `handlers`: Service Input and Output
* `plugins`: System plugins

### Goals
* Add Games
    * Riichi
        * make rules and validators
        * make simple input and output
        * make simple gameloop
        * add ai
* Add adapters
    * twitter adapter
    * irc adapter
    * twitch adapter
    * ncurses adapter
