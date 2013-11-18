#!/usr/bin/php
<?php // That top line tells php that this is a shell script. It's optional.
// Also note that mine is a custom setup at: /usr/local/bin/php, but most
// users probably have their php binary at the location above.

// Turning off all error reporting for this script. This is necessary b/c
// errors get printed to output, and output must be nothing but the
// alfred standard xml otherwise it doesn't work.
// error_reporting(0);


require_once 'workflows.php'; // Bring in WorkFlows class
$w = new Workflows;

// A note on $argv: the position of the args string in the array can differ
// depending on how the php binary is called. For instance:
//
// php -f filename.php -- "a string argument"
// ==> That command in shell will put the passed string at $argv[1].
//
// php filename.php -- "a string argument"
// ==> This puts it at $argv[2] b/c 'filename.php' is also passed as an argument.
$filter = false;
if ($argc > 0) {
  $filter = strtolower($argv[1]);
}

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
