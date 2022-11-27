package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

type Conn struct {
	*sqlx.DB
}

type Product struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

func (p *Product) Scan() {

}

func getConnString() string {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		connString = "user=postgres dbname=postgres sslmode=disable"
	}
	return connString
}

func CreateConnection(ctx context.Context) (*Conn, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", getConnString())
	if err != nil {
		return nil, err
	}
	return &Conn{
		DB: db,
	}, nil
}

func (c *Conn) QueryProducts(ctx context.Context, name string) ([]*Product, error) {
	products := []*Product{}
	err := c.SelectContext(ctx, &products, fmt.Sprintf("SELECT * FROM products WHERE name = '%s'", name))
	if len(products) == 0 {
		return nil, nil
	}
	return products, err
}
