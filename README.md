# Alfred Maestro

#### Activate any of your Keyboard Maestro macros in Alfred 2

This is a workflow for anyone who uses Keyboard Maestro and wishes it had built-in Alfred support. In Alfred simply type `km` followed by the name of any of your KM macros.

## Requirements

Alfred 2 (w/ Powerpack) and Keyboard Maestro 6.3 or greater. If you don't have them, go get them immediately. You will not regret it.

- [Alfred 2](http://www.alfredapp.com/) (Free, but requires £17 Powerpack)
- [Keyboard Maestro](http://www.keyboardmaestro.com/main/) (Free to try. $36/license)

Both well worth the cost.

## Installation

Download the zip file or clone this repo, then double-click the included 'AlfredMaestro.alfredworkflow' file to install.

## Usage

Type `km` followed by the name of any of your defined macros.

I also wrote a blog post on how I personally use this workflow for anyone who's interested: [Integrating Alfred and Keyboard Maestro][blogpost]

[blogpost]: http://blog.iansinnott.com/integrating-alfred-and-keyboard-maestro/

![usage example screen](screen.png "Usage Example")

#### Update:

Thanks to Peter Lewis of [Stairways Software][stair], the creator of Keyboard Maestro for adding a feature to KM that greatly increased the power of this workflow. Now Alfred can launch any of your macros, not just the ones with a hotkey!

[stair]: http://www.stairways.com/main/

## Possible Issues

### Application Specific Macro Groups

**Important:** If you have macros that you only want to run in certain applications read this.

If you create a macro group that is _not_ set to run in all applications you will need to make sure it is set to run in Alfred 2 in addition to whatever application you want it to run in. See [issue 5][issue5] for details.

[issue5]: https://github.com/iansinnott/alfred-maestro/issues/5

### PHP Version

This should run successfully with PHP 5.3 and above, which is standard on most macs. I think newer macs actually ship with 5.4. If you are getting empty results within Alfred then you may have an issue with an outdated version of PHP.

To figure out what version of PHP you are running you can try the following command in the Terminal: `php -v`. You can also run `which php` to see which binary is being used for the command line and where it's located.

## Troubleshooting

If you are having trouble, feel free to [open an issue][issues]. I try to check up on it regularly. Alternatively, ping me on Twitter: [@ian_989](https://twitter.com/ian_989)

[issues]: https://github.com/iansinnott/alfred-maestro/issues
