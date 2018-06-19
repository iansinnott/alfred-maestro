package main

import (
	"testing"
)

func TestGetKmMacros(t *testing.T) {
	expectedMacros := map[string]KmMacro{
		"UID-MACRO-1": KmMacro{"UID-MACRO-1", "Macro #1", "Category #1", ""},
		"UID-MACRO-2": KmMacro{"UID-MACRO-2", "Macro #2", "Category #1", "⌃⌥⌘X"},
		"UID-MACRO-3": KmMacro{"UID-MACRO-3", "Macro #3", "Category #2", "⌃⌥⌘R"},
	}

	actualMacros, err := getKmMacros()

	assertEqual(t, err, nil)
	assertEqual(t, len(expectedMacros), len(actualMacros))
	for uid, expectedMacro := range expectedMacros {
		assertEqual(t, expectedMacro, actualMacros[uid])
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
