package main

import (
	"encoding/xml"
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
	uid      string
	name     string
	category string
}

func getKmMacros() []KmMacro {
	command := os.Getenv("KM_MACROS_FETCH_COMMAND")
	if command == "" {
		command = "osascript ./km.scpt"
	}

	out, err := exec.Command("sh", "-c", command).Output()

	if err != nil {
		wf.Fatal("Unable to get macros from Keyboard Maestro")
		return nil
	}

	if !strings.Contains(string(out), "<?xml") {
		wf.Fatal(string(out))
		return nil
	}

	var categories KmCategories
	err = xml.Unmarshal(out, &categories)
	if err != nil {
		wf.Fatal("Unable to get macros from Keyboard Maestro")
		return nil
	}

	var macros []KmMacro

	for _, category := range categories.Categories {
		for _, item := range category.Items {
			macros = append(macros, KmMacro{
				uid:      item.getValueByKey("uid"),
				name:     item.getValueByKey("name"),
				category: category.getValueByKey("name"),
			})
		}
	}

	return macros
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
