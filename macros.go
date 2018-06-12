package main

import (
	"encoding/xml"
	"errors"
	"os"
	"os/exec"
	"strings"
)

type KmItem struct {
	Keys   []string `xml:"key"`
	Values []string `xml:"string"`
}

type KmCategory struct {
	Keys   []string `xml:"key"`
	Values []string `xml:"string"`
	Items  []KmItem `xml:"array>dict"`
}

type KmCategories struct {
	Categories []KmCategory `xml:"array>dict"`
}

type KmMacro struct {
	UID      string
	Name     string
	Category string
}

func getKmMacros() ([]KmMacro, error) {
	// Allow to change the command for fetching macros, so the function could be unit-tested
	command := os.Getenv("KM_MACROS_FETCH_COMMAND")
	if command == "" {
		command = "osascript ./km.scpt"
	}

	out, err := exec.Command("sh", "-c", command).Output()

	if err != nil {
		return nil, errors.New("Unable to get macros from Keyboard Maestro")
	}

	if !strings.Contains(string(out), "<?xml") {
		return nil, errors.New(string(out))
	}

	var categories KmCategories
	err = xml.Unmarshal(out, &categories)
	if err != nil {
		return nil, errors.New("Unable to get macros from Keyboard Maestro")
	}

	var macros []KmMacro

	for _, category := range categories.Categories {
		for _, item := range category.Items {
			macros = append(macros, KmMacro{
				UID:      item.getValueByKey("uid"),
				Name:     item.getValueByKey("name"),
				Category: category.getValueByKey("name"),
			})
		}
	}

	return macros, nil
}

func (item KmItem) getValueByKey(requestedKey string) string {
	for i, key := range item.Keys {
		if key == requestedKey {
			return item.Values[i]
		}
	}

	return ""
}

func (category KmCategory) getValueByKey(requestedKey string) string {
	isMacrosVisited := false
	for i, key := range category.Keys {
		if key == "macros" {
			isMacrosVisited = true
			if key == requestedKey {
				break
			}
		}

		if key == requestedKey {
			if isMacrosVisited {
				i--
			}

			return category.Values[i]
		}
	}

	return ""
}
