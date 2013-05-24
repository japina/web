package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"blackfriday"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var css string
	var repeat int
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	htmlFlags := blackfriday.HTML_COMPLETE_PAGE
	var renderer blackfriday.Renderer
	renderer = blackfriday.HtmlRenderer(htmlFlags, title, css)	

	title := "test"	
	filename := title+".txt"
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(w,"%s",err)
	}
	repeat = 1
	var body []byte
	for i := 0; i < repeat; i++ {
		body = blackfriday.Markdown(input, renderer, extensions)
	}

	fmt.Fprintf(w,"<h1>%s</h1> %s",title,string(body))
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}