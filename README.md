# Keyboard Maestro Workflow

#### Activate any of your KM hotkeys in Alfred 2

This is a workflow for anyone who uses Keyboard Maestro and wishes it had built-in Alfred support. In Alfred simply type `km` followed by the name of any of your KM hotkey macros.

## Installation

Download the zip file or clone this repo, then double-click the included 'KeyboardMaestro.alfredworkflow' file to install. 

## Usage

Type `km` followed by the name of any of your defined macros. 

![usage example screen](screen.png "Usage Example")

## Limitations
The interface that allows retreival of available macros only lists macros that have been assigned hotkeys or string triggers. So you may notice that if you have other macros that are triggered by say the pallete entry, they will not show up. Thus, for macros you would like to call by name through alfred I would suggest assigning them either an obscure key combination or text string that won't interfere with your other system hotkeys. 

## Possible Issues 
This runs successfully with PHP 5.5.5 (The latest version as of this writing). The XML used to generate the feedback for Alfred was having some issues on a lower version of PHP (5.3.26) when I tested. If you are getting empty results within Alfred then this may be the issue. 

To figure out what version of PHP you are running you can try the following command in the Terminal: `/usr/bin/php -v`. You can also run `which php` to see which binary is being used for the command line. 

