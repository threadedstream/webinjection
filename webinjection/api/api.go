package api

import (
	"context"
	"errors"
	"github.com/threadedstream/webinjection/webinjection/database"
	"net/http"
)

var (
	conn *database.Conn
)

func init() {
	var err error
	conn, err = database.CreateConnection(context.Background())
	if err != nil {
		panic(err)
	}
}

func FetchProducts(_ http.ResponseWriter, request *http.Request) ([]*database.Product, error) {
	if request.Method != http.MethodGet {
		return nil, errors.New("invalid method used")
	}
	name := request.FormValue("product_name")
	products, err := conn.QueryProducts(context.Background(), name)
	if err != nil {
		return nil, err
	}
	return products, err
}
