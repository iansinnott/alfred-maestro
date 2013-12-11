#!/usr/bin/php
<?php // The shebang probably not necessarr, but i'm leaving it in. 

// Turning off all error reporting for this script. This is necessary b/c
// errors get printed to output, and output must be nothing but the
// alfred standard xml otherwise it doesn't work.
error_reporting(0);

require_once 'workflows.php'; // Bring in WorkFlows class
$w = new Workflows;

// A note on $argv: the position of the args string in the array can differ
// depending on how the php binary is called. For instance:
//
// php -f filename.php -- "a string argument"
// ==> That command in shell will put the passed string at $argv[1].
//     $argv[0] holds the name of the file.
//
// php filename.php -- "a string argument"
// ==> This puts it at $argv[2] b/c 'filename.php' and '--' are also passed as arguments.
$filter = false;

// Make sure we have an arg to search for. When argc == 1 it just shows that 
// the first arg is the name of the file ('main.php'), which means there is still
// no argument. If we pass something like php -f main.php -- "" that empty string
// will register as an arg at argv[1], thus I must handle that case.


// This should probably be one statement, will worry about that later.
if ($argc <= 1) :
  no_result('Start typing a macro name...', 'No string was given.');
elseif (!$argv[1]) :
  no_result('Start typing a macro name...', 'Just type out the name of a macro.');
endif;

$filter = strtolower($argv[1]);

/**
 * Function to call if no result is found, or if they enter nothing.
 * @return xml
 */
function no_result($message = '', $submessage = '') {
  global $w; // Need this to output the messages

  // Defaults
  if (!$message) $message = 'No results found.';
  if (!$submessage) $submessage = 'No macros matched your query.';

  $w->result( 'reeder.result', 'na', $message, $submessage, 'icon.png', 'no' );
  echo $w->toXML();
  exit;
}

/**
 * This is the meat of the workflow. This takes the unformatted xml from alfred and
 * parses it into something we can actually use. 
 * @return array of arrays of macros
 */
function parse_xml() {
  $output = array();
  exec('osascript ./km.scpt', $output);
    if (!$output) return false;

  // For some reason this comes out as an array of 'lines' of xml.
  $output = join('', $output);

  $xml = new SimpleXMLElement($output);
  $key_groups = $xml->array->dict;
  $items = $xml->array->dict;

  $macros = array();

  foreach ($items as $item){
    foreach ($item->array->dict as $entry){
      $macro = array();
      // Originally had (array)$entry->string; but that didn't work 
      // for some reason in php 5.3.26. Works fine in 5.5.5 though. 
      // But this workflow is meant to be for everyone so I can hardly
      // ignore the default version of php that ships with macs.
      $entry = (array)$entry; // This is the extra step now
      $props = $entry['string'];
      if ($entry['key'][0] == 'key') {
        $macro['uid'] = $props[0];
        $macro['name'] = $props[1];
        $macro['sort'] = $props[2];
        // $macro['uid'] = $props[3];
        $macro['match'] = strtolower($macro['name']);
      } else {
        $macro['uid'] = $props[2];
        $macro['name'] = $props[0];
        $macro['sort'] = $props[1];
        // $macro['uid'] = $props[3];
        $macro['match'] = strtolower($macro['name']);
      } // endif
      $macros[] = $macro;
    }// endforeach
  } // endforeach


  return $macros;
}

$macros = parse_xml();
if (!$macros) :
  $w->result( 'reeder.result',
    'na',
    'No results found',
    'No search results found matching your query',
    'icon.png',
    'no'
  );
elseif ($filter) : // If keyword is given, do a filter
  foreach ($macros as $macro) {
    // Very simple filtering, but works quite well.
    // Real fuzzy matching would be nice, but thats a wishlist item.
    $match = strpos($macro['match'], $filter);
    if ($match !== false) :
      $w->result(
        $macro['uid'],
        $macro['uid'],
        $macro['name'],
        $macro['name'],
        'icon.png',
        'yes'
      );
    endif;
  }
  // If no matches, return this.
  if (!$w->results()) :
      $w->result( 'reeder.result', 'na', 'No results found', 'No macros matched your query', 'icon.png', 'no' );
  endif;
else : // Return all available macros
  foreach ($macros as $macro) {
    $w->result(
      $macro['uid'],
      $macro['uid'],
      $macro['name'],
      $macro['name'],
      'icon.png',
      'yes'
    );
  }
endif;

echo $w->toXML();
