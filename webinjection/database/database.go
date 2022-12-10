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

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func getConnString() string {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		connString = "host=postgres-db user=postgres dbname=postgres sslmode=disable"
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
	err := c.SelectContext(ctx, &products, fmt.Sprintf("SELECT * FROM products WHERE name = '%s';", name))
	if len(products) == 0 {
		return nil, nil
	}
	return products, err
}

func (c *Conn) QueryUser(ctx context.Context, username, password string) ([]*User, error) {
	users := []*User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE username = $1 AND password = '%s'", password)
	err := c.SelectContext(ctx, &users, query, username)
	if len(users) == 0 {
		return nil, nil
	}
	return users, err
}

func (c *Conn) QueryUserProtected(ctx context.Context, username, password string) ([]*User, error) {
	users := []*User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE username = $1 AND password = '%s'", password)
	err := c.SelectContext(ctx, &users, query, username)
	if len(users) == 0 {
		return nil, nil
	}
	return users, err
}
