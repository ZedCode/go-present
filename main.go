package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	presentation = flag.String("presentation", "example.md", "The presentation to inject into the template")
	theme        = flag.String("theme", "white", "The theme to apply to the presentation")
)

func init() {
	flag.StringVar(presentation, "p", "example.md", "The presentation to inject into the template")
	flag.StringVar(theme, "t", "white", "The theme to apply to the presentation")
}

func main() {
	flag.Parse()
	tc := TemplateConfig{MarkdownFile: *presentation, ThemeName: *theme}
	err := copyFile(filepath.Join("/root", "md", tc.MarkdownFile), filepath.Join("/root", "reveal.js", tc.MarkdownFile))
	if err != nil {
		log.Fatal(err)
	}
	t, err := template.ParseFiles("/root/md/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("/root/reveal.js/index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = t.Execute(f, tc)
	if err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir("/root/reveal.js"))
	http.Handle("/", http.StripPrefix("/", fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
