package models

type DemoItem struct {
	ID   int64  `db:"id, primarykey, autoincrement" json:"id"`
	Name string `db:"name" json:"name"`
}
