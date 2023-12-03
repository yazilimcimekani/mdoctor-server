package controllers

import (
	"fmt"
	"net/http"

	"github.com/yazilimcimekani/mdoctor-server/pkg/mdtohtml"
)

func Markdown() func(w http.ResponseWriter, r *http.Request) {
	const mdTemplatePath = "views/md.html"
	var html string = mdtohtml.LoadFile(mdTemplatePath)
	fullHTML := fmt.Sprintf(html, mdtohtml.MarkdownToHtml("# Hello World\nSa\n# Hello\n###### Hellooooooooo i am a h6\nAnd i'm a normal text\n"))

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fullHTML))
	}
}
