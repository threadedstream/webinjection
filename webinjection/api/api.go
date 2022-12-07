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

func Login(request *http.Request) (bool, error) {
	username := request.FormValue("username")
	password := request.FormValue("password")
	if username == "" && password == "" {
		return false, errors.New("username and password must be present")
	}
	users, err := conn.QueryUserProtected(context.Background(), username, password)
	if err != nil {
		return false, err
	}
	if users == nil {
		return false, nil
	}
	return true, nil
}
