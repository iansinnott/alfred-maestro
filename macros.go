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
	Values []string `xml:",any"`
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
	Hotkey   string
}

func getKmMacros() (map[string]KmMacro, error) {
	// Allow to change the command for fetching macros, so the function could be unit-tested
	getAllMacrosCommand := os.Getenv("GET_ALL_KM_MACROS_COMMAND")
	if getAllMacrosCommand == "" {
		getAllMacrosCommand = "osascript ./get_all_km_macros.scpt"
	}

	categoriesWithAllMacros, err := getKmCategories(getAllMacrosCommand)
	if err != nil {
		return nil, err
	}

	getHotkeyMacrosCommand := os.Getenv("GET_HOTKEY_KM_MACROS_COMMAND")
	if getHotkeyMacrosCommand == "" {
		getHotkeyMacrosCommand = "osascript ./get_hotkey_km_macros.scpt"
	}

	categoriesWithHotKeyMacros, err := getKmCategories(getHotkeyMacrosCommand)
	if err != nil {
		return nil, err
	}

	macros := make(map[string]KmMacro)
	var uid string

	for _, category := range categoriesWithAllMacros.Categories {
		for _, item := range category.Items {
			uid = item.getValueByKey("uid")
			macros[uid] = KmMacro{
				UID:      uid,
				Name:     item.getValueByKey("name"),
				Category: category.getValueByKey("name"),
				Hotkey:   "",
			}
		}
	}

	for _, category := range categoriesWithHotKeyMacros.Categories {
		for _, item := range category.Items {
			uid = item.getValueByKey("uid")
			macro, isExists := macros[uid]
			if isExists == true {
				macro.Hotkey = item.getValueByKey("key")
				// TODO Use pointer instead?
				macros[uid] = macro
			}
		}
	}

	return macros, nil
}

func getKmCategories(command string) (KmCategories, error) {
	out, err := exec.Command("sh", "-c", command).Output()

	var categories KmCategories
	if err != nil {
		return categories, errors.New("Unable to get macros from Keyboard Maestro")
	}

	if !strings.Contains(string(out), "<?xml") {
		return categories, errors.New(string(out))
	}

	err = xml.Unmarshal(out, &categories)
	if err != nil {
		return categories, errors.New("Unable to get macros from Keyboard Maestro")
	}

	return categories, nil
}

func (item KmItem) getValueByKey(requestedKey string) string {
	for i, key := range item.Keys {
		if key == requestedKey {
			return item.Values[i]
		}
	}

	return ""
}

// TODO Find out how to use the same func for both KmItem and KmCategory
func (item KmCategory) getValueByKey(requestedKey string) string {
	for i, key := range item.Keys {
		if key == requestedKey {
			return item.Values[i]
		}
	}

	return ""
}
