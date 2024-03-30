package entity

type Color struct {
	ID      int    `db:"id"`
	Name    string `db:"name"`
	HexCode string `db:"hex_code"`
}
