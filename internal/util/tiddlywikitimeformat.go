package util

import (
	"strings"
)

type tiddlyDateFormat = string
type goTimeFormat = string

var conversionMap = map[tiddlyDateFormat]goTimeFormat{
	"DDD":   "Monday",
	"ddd":   "Mon",
	"DD":    "2",
	"0DD":   "02",
	"DDth":  "",
	"WW":    "",
	"0WW":   "",
	"MMM":   "January",
	"mmm":   "Jan",
	"MM":    "1",
	"0MM":   "01",
	"YYYY":  "2006",
	"YY":    "06",
	"hh":    "15",
	"0hh":   "",
	"hh12":  "3",
	"0hh12": "03",
	"mm":    "4",
	"0mm":   "04",
	"ss":    "5",
	"0ss":   "05",
	"XXX":   "",
	"0XXX":  "",
	"am":    "pm",
	"pm":    "pm",
	"AM":    "PM",
	"PM":    "PM",
	"TZD":   "-0700",
	"[UTC]": "",
}

// FindIncompatibleTiddlyFormats returns a slice of the
// tiddly times not supported by this app (go)
func FindIncompatibleTiddlyFormats(tiddlytime string) []string {
	incompats := []string{}

	for tidformat, goformat := range conversionMap {
		if goformat == "" {
			incompats = append(incompats, tidformat)
		}
	}
	return incompats
}

// ConvertTiddlyTimeToGo replaces the tiddly time format
// codes with the go format codes
func ConvertTiddlyTimeToGo(tiddlytime string) string {

	gotime := tiddlytime
	for tidformat, goformat := range conversionMap {
		if goformat != "" {

			gotime = strings.ReplaceAll(gotime, tidformat, goformat)
		}
	}
	return gotime
}
