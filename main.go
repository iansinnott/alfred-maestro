package main

import (
	"time"

	"github.com/deanishe/awgo"
)

var wf *aw.Workflow

func run() {
	var kmMacroErr error
	reload := func() (interface{}, error) {
		macrosIndex, err := getKmMacros()

		if err != nil {
			kmMacroErr = err
		}

		var macros []KmMacro
		for _, macro := range macrosIndex {
			macros = append(macros, macro)
		}

		return macros, err
	}

	// Cache KM macros for 15 seconds
	maxCache := 15 * time.Second
	var macros []KmMacro
	err := wf.Cache.LoadOrStoreJSON("kmMacros", maxCache, reload, &macros)

	if err != nil {
		// LoadOrStoreJSON() generates a new error message
		// Therefore use kmMacroErr to get the original error message
		if kmMacroErr == nil {
			wf.Fatal(err.Error())
		} else {
			wf.Fatal(kmMacroErr.Error())
		}

		return
	}

	var item *aw.Item
	for _, macro := range macros {
		item = wf.NewItem(macro.Name).UID(macro.UID).Valid(true).Arg(macro.UID)
		if macro.Hotkey != "" {
			item.Subtitle(macro.Hotkey)
		}
	}

	args := wf.Args()
	var searchQuery string
	if len(args) > 0 {
		searchQuery = args[0]
	}

	if searchQuery == "" {
		wf.WarnEmpty("No macros found", "It seems that you haven't created any macros yet.")
	} else {
		wf.Filter(searchQuery)
		wf.WarnEmpty("No macros found", "Try a different query.")
	}

	wf.SendFeedback()
}

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(run)
}
