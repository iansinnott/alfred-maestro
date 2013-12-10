# Alfred Maestro

#### Activate any of your Keyboard Maestro macros in Alfred 2

This is a workflow for anyone who uses Keyboard Maestro and wishes it had built-in Alfred support. In Alfred simply type `km` followed by the name of any of your KM hotkey macros.

## Requirements 

Alfred 2 (w/ Powerpack) and Keyboard Maestro 6.3 or greater. If you don't have them, go get them immediately. You will not regret it.

- [Alfred 2](http://www.alfredapp.com/) (Free, but requires £17 Powerpack)
- [Keyboard Maestro](http://www.keyboardmaestro.com/main/) (Free to try. $36/license)

Both well worth the cost.

You will also need a relatively recent version of PHP available for command-line use. For anyone running **Mavericks** this shouldn't be a problem. Anyone who's not on Mavericks can run `php -v` within Terminal to see what version they have. We the 'Possible Issues' section bellow for more info. 

## Installation

Download the zip file or clone this repo, then double-click the included 'AlfredMaestro.alfredworkflow' file to install. 

## Usage

Type `km` followed by the name of any of your defined macros. 

![usage example screen](screen.png "Usage Example")

#### Update:

Thanks to Peter Lewis of [Stairways Software][stair], the creator of Keyboard Maestro for adding a feature to KM that greatly increased the power of this workflow. Now Alfred can launch any of your macros, not just the ones with a hotkey!

[stair]: http://www.stairways.com/main/

## Possible Issues 
This should run successfully with PHP 5.3 and above, which is standard on most macs. I think newer macs actually ship with 5.4. If you are getting empty results within Alfred then you may have an issue with an outdated version of PHP. 

To figure out what version of PHP you are running you can try the following command in the Terminal: `php -v`. You can also run `which php` to see which binary is being used for the command line and where it's located.

If you are having trouble, feel free to open an issue on this repo. I try to check up on it regularly. Alternatively, ping me on Twitter: [@ian_989](https://twitter.com/ian_989)

