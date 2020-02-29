package models

type UserModel struct {
	ID    int64  `db:"id, primarykey, autoincrement, omitempty" json:"id"`
	Name  string `db:"name, omitempty" json:"name"`
	Score int    `db:"score, omitempty" json:"score"`
}
