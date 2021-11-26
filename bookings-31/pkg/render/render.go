package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("RenderTemplate>", err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

}

// CreateTemplateCache creates a template cache as a map, page + base
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	// 폴더이름+파일명 저장
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// 폴더정보를 제거하고 파일 이름만 저장
		name := filepath.Base(page)

		// 페이지 정보 로딩
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		// for i, layout := range layouts {
		// 	log.Printf("layout(%d): %s", i, layout)
		// }
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			// 페이지 정보에 base 정보 결함.
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			// log.Printf("ts2: %v\n", *ts)
			if err != nil {
				return myCache, err
			}

		}
		myCache[name] = ts

	}

	return myCache, nil
}
