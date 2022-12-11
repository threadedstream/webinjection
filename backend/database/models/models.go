package models

type Product struct {
	ID          int    `db:"id" json:"id,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
	Description string `db:"description" json:"description,omitempty"`
}

type User struct {
	ID       int    `db:"id" json:"id,omitempty"`
	Username string `db:"username" json:"username,omitempty"`
	Password string `db:"password" json:"password,omitempty"`
}
