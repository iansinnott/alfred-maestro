package main

import (
	"testing"
)

func TestGetKmMacros(t *testing.T) {
	expectedMacros := [3]KmMacro{
		KmMacro{"UID-MACRO-1", "Macro #1", "Category #1"},
		KmMacro{"UID-MACRO-2", "Macro #2", "Category #1"},
		KmMacro{"UID-MACRO-3", "Macro #3", "Category #2"},
	}

	actualMacros, err := getKmMacros()

	assertEqual(t, err, nil)
	assertEqual(t, len(expectedMacros), len(actualMacros))
	for i, expectedMacro := range expectedMacros {
		assertEqual(t, expectedMacro, actualMacros[i])
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
