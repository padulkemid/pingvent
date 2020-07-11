// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Barang struct {
	ID        string  `json:"id"`
	Nama      string  `sql:",unique" json:"nama"`
	Harga     float64 `json:"harga"`
	Stock     int     `json:"stock"`
	Vendor    string  `json:"vendor"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type BarangBaru struct {
	Nama   string  `json:"nama"`
	Harga  float64 `json:"harga"`
	Stock  int     `json:"stock"`
	Vendor string  `json:"vendor"`
}

type User struct {
	ID       string `json:"id"`
	Username string `sql:",unique" json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserBaru struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}