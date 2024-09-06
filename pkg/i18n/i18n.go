package i18n

import (
	"path/filepath"
	"strings"

	"github.com/kaptinlin/go-i18n"
)

var (
	Locales = GetCurrentLocales()
	I18n    *i18n.I18n
)

func init() {
	var withLocaleFuncs = make([]func(*i18n.I18n), len(Locales)+1)
	withLocaleFuncs[0] = i18n.WithDefaultLocale("en-US")

	// Load the locales handler
	for i, locale := range Locales {
		withLocaleFuncs[i+1] = i18n.WithLocales(locale)
	}
	I18n = i18n.NewBundle(withLocaleFuncs...)

	// Load the messages
	I18n.LoadGlob(filepath.Join("locales", "*.json"))
}

func GetCurrentLocales() []string {
	files, _ := filepath.Glob("locales/*.json")
	locales := make([]string, len(files))

	for i, f := range files {
		locales[i] = strings.TrimSuffix(filepath.Base(f), ".json")
	}

	return locales
}

func Get(locale string, key string, data ...i18n.Vars) string {
	return I18n.NewLocalizer(locale).Get(key, data...)
}

func GetCountryCode(locale string) string {
	var l = strings.Split(locale, "-")
	if len(l) == 1 {
		return l[0]
	}
	return l[1]
}
