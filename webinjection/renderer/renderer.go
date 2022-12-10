package renderer

import (
	"bytes"
	"fmt"
	"github.com/threadedstream/webinjection/webinjection/database"
	"html/template"
	"io"
	"os"
)

var (
	projectRoot string
)

func init() {
	// TODO(threadedstream): currently, it depends on the location of a binary
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	projectRoot = wd
}

func newTemplate(path, name string) (tmpl *template.Template, err error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	tmpl, err = template.New(name).Parse(string(contents))
	return tmpl, err
}

func RenderIndex() []byte {
	w := bytes.NewBuffer(nil)
	x := struct {
		Message string
	}{
		Message: "SQL Injection is no jokes",
	}

	t, err := newTemplate(fmt.Sprintf("%s/static/%s", projectRoot, "index.html"), "index.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, x)
	if err != nil {
		panic(err)
	}
	return w.Bytes()
}

func RenderProductInfo(products []*database.Product) []byte {
	w := bytes.NewBuffer(nil)
	x := struct {
		Products []*database.Product
	}{
		Products: products,
	}

	t, err := newTemplate(fmt.Sprintf("%s/static/%s", projectRoot, "product_info.html"), "product_info.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, x)
	if err != nil {
		panic(err)
	}
	return w.Bytes()
}

func RenderStatic(name string) []byte {
	f, err := os.OpenFile(fmt.Sprintf("%s/static/%s", projectRoot, name), os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return bs
}
