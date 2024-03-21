package model

import "time"

type Goods struct {
	ID         int
	CategryID  int
	Name       string
	Removed    bool
	Created_at time.Time
}

type Categry struct {
	ID   int
	Name string
}
