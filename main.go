package main

import (
	"github.com/deanishe/awgo"
)

var wf *aw.Workflow

func run() {
	macros := getKmMacros()
	if macros == nil {
		return
	}

	for _, macro := range macros {
		wf.NewItem(macro.name).UID(macro.uid).Valid(true).Arg(macro.uid)
	}

	args := wf.Args()
	var searchQuery string
	if len(args) > 0 {
		searchQuery = args[0]
	}

	if searchQuery == "" {
		wf.WarnEmpty("No Keyboard Maestro macros found", "It seems that you haven't created any macro yet.")
	} else {
		wf.Filter(searchQuery)
		wf.WarnEmpty("No Keyboard Maestro macros found that matched your query", "Try a different query?")
	}

	wf.SendFeedback()
}

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(run)
}
