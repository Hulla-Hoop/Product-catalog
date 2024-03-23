package model

import "time"

type Goods struct {
	ID         int
	Updated_at time.Time
	Name       string
	Removed    bool
	Created_at time.Time
}

type Category struct {
	ID         int
	Name       string
	Removed    bool
	Updated_at time.Time
	Created_at time.Time
}

type Relation struct {
	Goods_ID int
	Category int
	Removed  bool
}

type Product struct {
	Name     string
	Category string
}
