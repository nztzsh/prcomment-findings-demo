package main

import "golang.org/x/text/language"

func parseLang(tag string) string {
	t, err := language.Parse(tag)
	if err != nil {
		return "und"
	}
	return t.String()
}
