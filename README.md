<p align="center">
  <img src="https://4kbu.files.wordpress.com/2015/08/10_amae_koromo_3840x2160_.png" width="350"/>
</p>
<h1 align="center">Koromo</h1>
<p align="center">A bot javascript bot framework for Riichi Mahjong</p>

### Summary
Koromo is a general bot framework originally forked from [leafo's bot](https://github.com/leafo/saltw-bot) with architectural influence from [munar](https://github.com/welovekpop/munar). Koromo looks to provide a robust framework for plugins to interope connections with a variety of other chat services. The main intention for Koromo is a strong AI for Riichi Mahjong because I want another option from Tenhou.

### Features
Koromo is divided into three parts:
* `core`: an event loop that listens and outputs
* `adapters`: an implementation of a listener
* `plugins`: an implementation of an output

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
