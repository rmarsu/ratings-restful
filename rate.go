package todo

type Stars struct {
	Id    int    `json:"id" db:"id"`
	Obj   string `json:"object" db:"obj"`
	Stars int    `json:"stars" db:"stars"`
}
