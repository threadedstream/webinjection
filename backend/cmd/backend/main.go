package main

import (
	"encoding/json"
	"github.com/threadedstream/webinjection/backend/api"
	"github.com/threadedstream/webinjection/backend/renderer"

	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	doOnce sync.Once
)

var (
	addr = flag.String("addr", "0.0.0.0:8000", "address to bind server to")
)

func setupHandlers() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(renderer.RenderHomePage())
	})
	mux.HandleFunc("/sqli_1", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(renderer.RenderStatic("sqli_1/index.html"))
	})
	mux.HandleFunc("/product_info", func(writer http.ResponseWriter, request *http.Request) {
		products, err := api.FetchProducts(writer, request)
		if err != nil {
			log.Println(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Write(renderer.RenderProductInfo(products))
	})
	mux.HandleFunc("/sqli_2", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			writer.Write(renderer.RenderStatic("index.html"))
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

	mux.HandleFunc("/api/v1/products", func(writer http.ResponseWriter, request *http.Request) {
		products, err := api.GetAllProducts(context.Background())
		if err != nil {
			log.Println(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		bs, err := json.Marshal(products)
		if err != nil {
			log.Println(err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Write(bs)
	})
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(renderer.RenderStatic("not_found.html"))
	})
	return mux
}

func runHTTPServer(ctx context.Context) {
	doOnce.Do(func() {
		mux := setupHandlers()
		s := &http.Server{Addr: *addr, Handler: mux}

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

		log.Println("Running server on ", *addr)
		if err := s.ListenAndServe(); err != nil {
			panic(err)
		}
	})
}

func main() {
	ctx := context.Background()
	flag.Parse()
	runHTTPServer(ctx)
}
