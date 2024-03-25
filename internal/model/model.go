package model

import "time"

type Goods struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Removed    bool      `json:"removed"`
	Updated_at time.Time `json:"update_at"`
	Created_at time.Time `json:"created_at"`
}

type Category struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Removed    bool      `json:"removed"`
	Updated_at time.Time `json:"updated_at"`
	Created_at time.Time `json:"created_at"`
}

type Relation struct {
	Goods_ID int  `json:"goods_id"`
	Category int  `json:"category_id"`
	Removed  bool `json:"removed"`
}

type Product struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}
