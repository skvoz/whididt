package lib

import (
	"log"
	"path/filepath"
	"text/template"
)

func funcMap() template.FuncMap {
	funcMap := template.FuncMap{
		"loop": func(from, to int) <-chan int {
			ch := make(chan int)
			go func() {
				for i := from; i <= to; i++ {
					ch <- i
				}
				close(ch)
			}()
			return ch
		},
	}

	return funcMap
}

func render(fileName string, v ...map[string]string) *template.Template {

	tmpl, err := template.New("loopTest").Funcs(funcMap()).
		ParseFiles(filepath.Join("../../", fileName))
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	return tmpl
}
