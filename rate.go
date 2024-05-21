package todo

type Stars struct {
	Object string `json:"object" db:"object"`
	Stars  int    `json:"stars" db:"stars"`
}
