package model

import (
	"encoding/json"
	"fmt"
)

//https://www.alfredapp.com/help/workflows/inputs/script-filter/json/
type (
	Icon struct {
		Type string `json:"type,omitempty"`
		Path string `json:"path,omitempty"`
	}
	Mod  map[string]string
	Item struct {
		UID          string         `json:"uid,omitempty"`
		Type         string         `json:"type,omitempty"`
		Title        string         `json:"title,omitempty"`
		Subtitle     string         `json:"subtitle,omitempty"`
		Arg          string         `json:"arg,omitempty"`
		Autocomplete string         `json:"autocomplete,omitempty"`
		Icon         *Icon          `json:"icon,omitempty"`
		Valid        bool           `json:"valid,omitempty"`
		Match        string         `json:"match,omitempty"`
		Mods         map[string]Mod `json:"mods,omitempty"`
		Quicklookurl string         `json:"quicklookurl,omitempty"`
		Text         Mod            `json:"text,omitempty"`
	}

	Mgr struct {
		Items []*Item `json:"items,omitempty"`
	}
)

func (mgr *Mgr) AddItem(item *Item) {
	mgr.Items = append(mgr.Items, item)
}

func (item *Item) AddMod(cmd string, mod Mod) {
	if item.Mods == nil {
		item.Mods = make(map[string]Mod)
	}
	item.Mods[cmd] = mod
}

func (mgr *Mgr) PrintJson() {
	b, e := json.Marshal(mgr)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(b))
}

func (i *Icon) For(item *Item) *Item {
	if item != nil {
		item.Icon = i
	}
	return item
}
