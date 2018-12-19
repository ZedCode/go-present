package main

import (
	"io/ioutil"
)

// TemplateConfig is exported for parsing templates
// Config note, available templates are:
//  beige, black, blood, league, moon
//  night, serif, simple, sky, solarized
//    default is "white"
type TemplateConfig struct {
	MarkdownFile string // defaults to example.md
	ThemeName    string // optional
}

func copyFile(src string, dst string) error {
	d, e := ioutil.ReadFile(src)
	if e != nil {
		return e
	}
	e = ioutil.WriteFile(dst, d, 0644)
	if e != nil {
		return e
	}
	return nil
}
