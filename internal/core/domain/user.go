package domain

type User struct {
	Username  string `json:"username" validate:"required" bun:"notnull"`
	Password  string `json:"password" validate:"required" bun:"notnull"`
	Salt      string `bun:"notnull"`
	Icon      string `json:"icon"`
	Status    string `bun:"notnull"`
	Privilege string `bun:"notnull"`
	Id        int32  `json:"id" bun:"pk,autoincrement"`
}
