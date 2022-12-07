package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"github.com/threadedstream/webinjection/webinjection/api"
	"github.com/threadedstream/webinjection/webinjection/database"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	projectRoot string
	doOnce      sync.Once
)

func init() {
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

func renderIndex() []byte {
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

func renderProductInfo(products []*database.Product) []byte {
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

func renderStatic(name string) []byte {
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

func setupHandlers() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(renderIndex())
	})
	mux.HandleFunc("/level1", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(renderStatic("level1.html"))
	})
	mux.HandleFunc("/product_info", func(writer http.ResponseWriter, request *http.Request) {
		products, err := api.FetchProducts(writer, request)
		if err != nil {
			log.Println(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Write(renderProductInfo(products))
	})
	mux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			writer.Write(renderStatic("login.html"))
		} else if request.Method == "POST" {
			canBeLoggedIn, err := api.Login(request)
			if err != nil {
				log.Println(err.Error())
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			if canBeLoggedIn {
				http.Redirect(writer, request, "/home", http.StatusMovedPermanently)
				return
			}
			http.Redirect(writer, request, "/login?message=Authentication failed", http.StatusMovedPermanently)
		} else {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(renderStatic("not_found.html"))
	})
	return mux
}

func runHTTPServer(ctx context.Context) {
	doOnce.Do(func() {
		mux := setupHandlers()
		s := &http.Server{Addr: "0.0.0.0:8000", Handler: mux}

		notifier := make(chan os.Signal, 1)
		signal.Notify(notifier, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
		go func() {
			for {
				select {
				case <-notifier:
					println("Gracefully quitting...")
					s.Shutdown(ctx)
					os.Exit(1)
				default:
					continue
				}
			}
		}()

		go func() {
			for {
				if err := recover(); err != nil {
					fmt.Print(err)
					// do proper handling
					os.Exit(1)
				}
			}
		}()

		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	})
}

func main() {
	ctx := context.Background()
	runHTTPServer(ctx)
}
